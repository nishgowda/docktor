package healthcheck
 


import (

	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/filters"
	"strconv"
)

var ctx = context.Background();

// PerformHealthCheck below functions on all running containers
func PerformHealthCheck(params...string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}
	var filter filters.Args
	var containers []types.Container
	
	if (params != nil) {
		filter = filters.NewArgs()
		filter.Add("ID", params[0])
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{Filters:filter})
		if err != nil {
			panic(err)
		}
		for _, container := range containers {
			port := strconv.FormatUint(uint64(container.Ports[1].PublicPort), 10)
			KillContainer(container.ID[:10])
			CreateContainer(container.Image, container.Ports[1].IP, port)
		}
	}else{
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
		for _, container := range containers {
			//fmt.Println(container.Ports[0].PublicPort)
			port := strconv.FormatUint(uint64(container.Ports[0].PublicPort), 10)
			KillContainer(container.ID[:10])
			CreateContainer(container.Image, container.Ports[0].IP, port)
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
// CreateContainer adds health check to existing contianer and restarts it
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
		panic("Unable to get the port")
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
		panic(err)
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return nil
}