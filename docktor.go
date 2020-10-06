package docktor


import(
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"strconv"
)

var ctx = context.Background();

// GetContainers performs below functions on all running containers
func GetContainers() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		port := strconv.FormatUint(uint64(container.Ports[1].PublicPort), 10)
		KillContainer(container.ID[:10])
		CreateContainer(container.Image, container.Ports[1].IP, port)
	}
}
// KillContainer kills all exisiting contianer to later add the health checks
func KillContainer(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}
	errs := cli.ContainerKill(ctx, containerID, "SIGKILL");
	if errs != nil {
		panic(errs)
	}
	fmt.Println("Killed container")
}
// CreateContainer adds health check to existing contianer and restarts it
func CreateContainer(contianerImage string, hostIP string, port string) string {
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
			Test : []string{"CMD-SHELL", "curl localhost:3000"},
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
	return "Success"
}