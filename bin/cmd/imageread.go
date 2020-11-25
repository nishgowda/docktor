package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/suggestions"
)

func init() {
	rootCmd.AddCommand(suggestCmd)
}

var suggestCmd = &cobra.Command {
	Use: "suggest",
	Short: "Suggests possible improvements to be made in a Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		suggestions.ReadImage(file)
	},
}