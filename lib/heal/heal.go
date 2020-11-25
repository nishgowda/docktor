package heal

import (
	"os"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"time"
	"errors"
)

var ctx = context.Background()

// GetUnheatlhyContainers fetches the ids of the unhealthy containers
func GetUnheatlhyContainers(params...string) []string {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var containerIDs = []string{}
	var containers []types.Container
	if (params == nil) {
		containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
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
	if (len(containerIds) < 2) {
		for _, id := range GetUnheatlhyContainers() {
			containerIds = append(containerIds, id)
		}
	}
	if (len(containerIds) == 0) {
		return errors.New("No containers were running")
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())	
	if err != nil {
		return err
	}	
	var timeout *time.Duration
	for _, id := range containerIds {
		if (len(id) > 2) {
			timeoutValue := time.Duration(10) * time.Second
			timeout = &timeoutValue
			e := cli.ContainerRestart(ctx, id, timeout)
			if e != nil {
				return e
			}
			log.Printf("Restarted container: %s\n", id)
		}
	}
	return nil
}

