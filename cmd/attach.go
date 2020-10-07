package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/healthcheck"
)

func init(){
	rootCmd.AddCommand(attachCmd)
}

var attachCmd = &cobra.Command{
	Use: "attach",
	Short: "Execute health checks with given containers",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) < 1) {
			healthcheck.PerformHealthCheck()
		}else {

			healthcheck.PerformHealthCheck(args[0])
		}
	},
}