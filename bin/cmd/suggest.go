package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/suggestions"
	"log"
)

func init() {
	rootCmd.AddCommand(suggestCmd)
}

var suggestCmd = &cobra.Command {
	Use: "suggest",
	Short: "Suggest possible improvements to be made in a Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		err := suggestions.ReadImage(file)
		if err != nil {
			log.Fatal(err)
		}
	},
}