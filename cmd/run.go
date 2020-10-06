package cmd

import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib"
)

func init(){
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use: "run",
	Short: "execute health checks with given containers",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) < 1) {
			lib.Perform()
		}else {
			lib.Perform(args[0])
		}
	},
}