package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/heal"
)

func init() {
	rootCmd.AddCommand()
}
var healCmd = &cobra.Command{
	Use: "heal",
	Short: "Heals unhealthy containers",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) < 1){
			containers := heal.GetUnheatlhyContainers()
			heal.ContainerHeal(containers)
		} else {
			heal.ContainerHeal(args)
		}
	},
}