package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/scan"
	"fmt"
)

func init() {
	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use: "scan",
	Short: "Scan images for vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		out := scan.Vulnerabilities(image)
		if len(file) > 1 {
			scan.WriteFile(out, file)
		} else {
			fmt.Println(out)
		}
	},
}
