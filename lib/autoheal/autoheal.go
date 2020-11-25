package autoheal

import (
	"log"
	"fmt"
	"os/exec"
)


// AutoHeal wraps auto update on containers
func AutoHeal(params []string) error {
	app := "docker"
    arg0 := "update"
    arg1 := "--restart"
	arg2 := "unless-stopped"

	if len(params) < 1 {
		log.Fatal("No container specified\n")
	}
	for _, id := range params {
		fmt.Println(id)
		cmd := exec.Command(app, arg0, arg1, arg2, id)
		err := cmd.Run()
		if err != nil {
			return err
		}
		log.Printf("Succesfully added autoheals to container: %s\n", id);
	}
	return nil
 }