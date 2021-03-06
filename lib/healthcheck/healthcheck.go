// Package healthcheck provides primitives that allow users to attach
// health checks to running docker containers
package healthcheck

import (
	"context"
	"errors"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var ctx = context.Background()

// PerformHealthCheck adds health checks to running containers
func PerformHealthCheck(params []string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	var filter filters.Args
	var containers []types.Container
	msg := ""
	// checks two cases: passed in container ids or none
	if len(params[0]) > 0 {
		// specified container given in the parameter
		filter = filters.NewArgs()
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{Filters: filter})
		if err != nil {
			return "", err
		}
		for _, container := range containers {
			// check if container is running on a port
			if len(container.Ports) > 0 {
				port := strconv.FormatUint(uint64(container.Ports[0].PublicPort), 10)
				killContainer(container.ID[:10])
				createContainer(container.Image, container.Ports[0].IP, port)
				msg += "Successfully added health checks to the following container: " + container.ID
			} else {
				msg += "The following container did not have a published port" + container.ID
			}
		}
		return msg, nil
	}
	// no containers specified so find the containers running through Docker API
	containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", errors.New("No running container detected")
	}
	for _, container := range containers {
		// check if container is running on a port
		if len(container.Ports) > 0 {
			port := strconv.FormatUint(uint64(container.Ports[0].PublicPort), 10)
			killContainer(container.ID[:10])
			createContainer(container.Image, container.Ports[0].IP, port)
			msg += "Successfully added health checks to the following container: " + container.Image
		} else {
			msg += "The following container did not have a published port" + container.ID
		}
	}
	return msg, nil
}

// killContainer kills all exisiting container to later add the health checks
func killContainer(containerID string) error {
	if len(containerID) == 0 {
		return errors.New("Invalid Arguments specified")
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	err = cli.ContainerKill(ctx, containerID, "SIGKILL")
	if err != nil {
		return err
	}
	return nil
}

// createContainer recreates a container with a health check added to it
func createContainer(containerImage string, hostIP string, port string) error {
	if len(containerImage) == 0 || len(hostIP) == 0 || len(port) == 0 {
		return errors.New("Invalid Arguments")
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	hostBinding := nat.PortBinding{
		HostIP:   hostIP,
		HostPort: port,
	}
	containerPort, err := nat.NewPort("tcp", port)
	if err != nil {
		return err
	}
	// Provide the features ensuring the container can get a health check
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: containerImage,
		Healthcheck: &container.HealthConfig{
			Test:        []string{"CMD-SHELL", "curl localhost:" + port},
			Interval:    1000000000,
			Retries:     10,
			Timeout:     10000000000,
			StartPeriod: 60000000000,
		},
		ExposedPorts: nat.PortSet{
			containerPort: struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: portBinding,
	}, nil, nil, "")

	if err != nil {
		return err
	}
	// respin the container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}
