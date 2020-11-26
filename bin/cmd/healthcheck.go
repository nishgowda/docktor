package cmd


import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/healthcheck"
	"log"
)


func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

var healthCheckCmd = &cobra.Command{
	Use: "healthcheck",
	Short: "Attach health checks on containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := healthcheck.PerformHealthCheck(containers)
		if err != nil {
			log.Fatal(err)
		}
	},
}
