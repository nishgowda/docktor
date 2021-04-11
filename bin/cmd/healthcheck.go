package cmd

import (
	"log"

	"github.com/nishgowda/docktor/lib/healthcheck"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

var healthCheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Attach health checks on containers",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := healthcheck.PerformHealthCheck(containers)
		if err != nil {
			log.Fatal(err)
		}
		if len(msg) < 1 {
			log.Fatal("No running containers detected", err)
		}
		log.Println(msg)
	},
}
