package cmd

import (
	"log"

	"github.com/nishgowda/docktor/lib/suggestions"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(suggestCmd)
}

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest possible improvements to be made in a Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := suggestions.ReadImage(file)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(msg)
	},
}
