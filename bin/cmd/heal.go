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
	Short: "Heal unhealthy containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := heal.ContainerHeal(containers)
		if (err != nil) {
			log.Fatal(err)
		}
	},
}
