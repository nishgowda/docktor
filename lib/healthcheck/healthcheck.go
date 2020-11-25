package healthcheck
 


import (
	"context"
	"log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/filters"
	"strconv"
	"errors"
)

var ctx = context.Background();

// PerformHealthCheck adds health checks to running containers
func PerformHealthCheck(params []string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}
	var filter filters.Args
	var containers []types.Container
	
	if (len(params) > 1) {
		filter = filters.NewArgs()
		log.Println(params)
		filter.Add("ID", params[0])
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{Filters:filter})
		if err != nil {
			return err
		}
		for _, container := range containers {
			port := strconv.FormatUint(uint64(container.Ports[1].PublicPort), 10)
			KillContainer(container.ID[:10])
			CreateContainer(container.Image, container.Ports[1].IP, port)
		}
	} else {
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
		if len(containers) < 1 {
			return errors.New("No running containers detected")
		}
		for _, container := range containers {
			port := strconv.FormatUint(uint64(container.Ports[0].PublicPort), 10)
			KillContainer(container.ID[:10])
			CreateContainer(container.Image, container.Ports[0].IP, port)
			log.Printf("Succesfully added health checks to the following container: %s\n", container.Image)
		}
	}
	return nil
}

// KillContainer kills all exisiting contianer to later add the health checks
func KillContainer(containerID string) error{
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}
	errs := cli.ContainerKill(ctx, containerID, "SIGKILL");
	if errs != nil {
		panic(errs)
	}
	return nil
}

// CreateContainer recreates a contianer with a health check added to it
func CreateContainer(contianerImage string, hostIP string, port string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}

	hostBinding := nat.PortBinding{
		HostIP:   hostIP,
		HostPort: port,
	}
	containerPort, err := nat.NewPort("tcp", port)
	if err != nil {
		return errors.New("Unable to get the port of container")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
		resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: contianerImage,
		Healthcheck: &container.HealthConfig {
			Test : []string{"CMD-SHELL", "curl localhost:" + port},
			Interval : 1000000000,
			Retries : 10,
			Timeout : 10000000000,
			StartPeriod : 60000000000,
		},
		ExposedPorts: nat.PortSet{
			containerPort: struct{}{},
		},
		}, &container.HostConfig{
			PortBindings: portBinding, 
		},nil, nil,"")
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}