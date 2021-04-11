package cmd

import (
	"log"

	"github.com/nishgowda/docktor/lib/scan"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan images for vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := scan.Vulnerabilities(image)
		if err != nil {
			log.Fatal(err)
		}
		if len(file) > 1 {
			scan.WriteFile(out, file)
		} else {
			log.Println("Success")
		}
	},
}
