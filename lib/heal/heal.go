package heal

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"fmt"
	"time"
)

var ctx = context.Background()

// GetUnheatlhyContainers fetches the ids of the unhealthy containers
func GetUnheatlhyContainers(params...string) []string {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}
	var containerIDs = []string{}
	var containers []types.Container
	if (params == nil) {
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
		for _, container := range containers {
			containerIDs = append(containerIDs, container.ID[:10])
		}
	} else {
		containerIDs = append(containerIDs, params[0])	
	}
	return containerIDs
}

// ContainerHeal heals unhealthy containers
func ContainerHeal(containerIds []string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		panic(err)
	}	
	var timeout *time.Duration
	for _, id := range containerIds {
		timeoutValue := time.Duration(10) * time.Second
		timeout = &timeoutValue
		e := cli.ContainerRestart(ctx, id, timeout)
		if e != nil {
			return e
		}
		fmt.Println("Restarted container: ", id)
	}
	return nil
}

