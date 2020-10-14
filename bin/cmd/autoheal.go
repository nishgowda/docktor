package cmd


import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/autoheal"
)
func init() {
	rootCmd.AddCommand(autoHealCmd)
}
var autoHealCmd = &cobra.Command{
	Use: "autoheal",
	Short: "Auto heal containers",
	Run: func(cmd *cobra.Command, args []string) {
		autoheal.AutoHeal(containers)
	},
}
