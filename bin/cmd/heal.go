package cmd

import (
	"github.com/nishgowda/docktor/lib/heal"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(healCmd)
}

var healCmd = &cobra.Command{
	Use:   "heal",
	Short: "Heal unhealthy containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := heal.ContainerHeal(containers)
		if err != nil {
			log.Fatal(err)
		}
	},
}
