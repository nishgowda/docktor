package cmd

import (
	"log"

	"github.com/nishgowda/docktor/lib/autoheal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(autoHealCmd)
}

var autoHealCmd = &cobra.Command{
	Use:   "autoheal",
	Short: "Add automatic healing to containers",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := autoheal.AutoHeal(containers)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(msg)
	},
}
