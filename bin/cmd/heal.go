package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/heal"

)
func init() {
	rootCmd.AddCommand(healCmd)
}

var healCmd = &cobra.Command{
	Use: "heal",
	Short: "Heals unhealthy containers",
	Run: func(cmd *cobra.Command, args []string) {
		if (containers == nil) {
			c := heal.GetUnheatlhyContainers()
			result:= heal.ContainerHeal(c)
			if (result != nil) {
				log.Fatal("There was an error in healing your container\n")
			}
		} else {
			result:= heal.ContainerHeal(containers)
			if (result != nil) {
				log.Fatal("There was an error in healing your container\n")
			}
		}
	},
}
