// Package autoheal provides wrapping on docker containers that restarts
// containers automatically when they become uhealthy
package autoheal

import (
	"errors"
	"os/exec"
)

// AutoHeal wraps auto restarts to containers when they become unhealthy
func AutoHeal(params []string) (string, error) {
	app := "docker"
	arg0 := "update"
	arg1 := "--restart"
	arg2 := "unless-stopped"
	arg3 := "ng"
	msg := ""
	if len(params) < 1 {
		return msg, errors.New("No container specified")
	}
	for _, id := range params {
		cmd := exec.Command(app, arg0, arg1, arg2, id, arg3)
		err := cmd.Run()
		if err != nil {
			return msg, err
		}
		msg += "Successfully added autoheals to container: " + id
	}
	return msg, nil
}
