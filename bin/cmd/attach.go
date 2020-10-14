package cmd


import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/healthcheck"
)


func init() {
	rootCmd.AddCommand(attachCmd)
}

var attachCmd = &cobra.Command{
	Use: "attach",
	Short: "Attach health checks on containers",
	Run: func(cmd *cobra.Command, args []string) {
		healthcheck.PerformHealthCheck(containers)
	},
}
