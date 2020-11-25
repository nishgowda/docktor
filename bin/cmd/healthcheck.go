package cmd


import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/healthcheck"
)


func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

var healthCheckCmd = &cobra.Command{
	Use: "healthcheck",
	Short: "Attach health checks on containers",
	Run: func(cmd *cobra.Command, args []string) {
		healthcheck.PerformHealthCheck(containers)
	},
}
