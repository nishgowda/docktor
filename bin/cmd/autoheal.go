package cmd


import (
	"github.com/spf13/cobra"
	"github.com/nishgowda/docktor/lib/autoheal"
	"log"
)

func init() {
	rootCmd.AddCommand(autoHealCmd)
}

var autoHealCmd = &cobra.Command{
	Use: "autoheal",
	Short: "Add automatic healing to containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := autoheal.AutoHeal(containers)
		if err != nil {
			log.Fatal(err)
		}
	},
}
