package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/autoheal"
)


var autoHealCmd = &cobra.Command {
	Use: "autoHeal",
	Short: "Automatically restart containers when unhealthy",
	Args: func (cmd *cobra.Command, args []string) error {
		if (len(args) < 1) {
			log.Fatal("Container id is required")
		}
		return nil
	},
	Run: func (cmd *cobra.Command, args []string) {
		autoheal.AutoHeal(args)
	},
}