// Package autoheal provides wrapping on docker containers that restarts
// containers automatically when they become uhealthy
package autoheal

import (
	"errors"
	"log"
	"os/exec"
)

// AutoHeal wraps auto restarts to containers when they become unhealthy
func AutoHeal(params []string) error {
	app := "docker"
	arg0 := "update"
	arg1 := "--restart"
	arg2 := "unless-stopped"

	if len(params) < 1 {
		return errors.New("No container specified")
	}
	for _, id := range params {
		cmd := exec.Command(app, arg0, arg1, arg2, id)
		err := cmd.Run()
		if err != nil {
			return err
		}
		log.Printf("Successfully added autoheals to container: %s\n", id)
	}
	return nil
}
