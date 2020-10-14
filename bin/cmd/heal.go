package cmd

import (
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
			if len(containers) < 1 {
				c := heal.GetUnheatlhyContainers()
				heal.ContainerHeal(c)
			} else {
				c := heal.GetUnheatlhyContainers(containers[0])
				heal.ContainerHeal(c)
			}
	},
}
