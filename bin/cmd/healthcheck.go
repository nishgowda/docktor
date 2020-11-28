package cmd

import (
	"github.com/nishgowda/docktor/lib/healthcheck"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

var healthCheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Attach health checks on containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := healthcheck.PerformHealthCheck(containers)
		if err != nil {
			log.Fatal(err)
		}
	},
}
