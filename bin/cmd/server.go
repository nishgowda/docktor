package cmd

import (
	"log"

	"github.com/nishgowda/docktor/bin/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command {
	Use: "server",
	Short: "Start docktor sever",
	Run: func(cmd *cobra.Command, args[] string) {
		log.Println("Running on localhost:" + port)
		server.Start(port)
	},
}
