package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

var (
	cfgFile string
	userLicense	string
	containers []string
	rootCmd = &cobra.Command{
	Use: "docktor",
	Short: "A health check generator for docker containers",
}
)
// Execute rootCmd
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Nish Gowda", "Nish Gowda")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "MIT")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Nish Gowda nish.gowda6@gmail.com")
	viper.SetDefault("license", "MIT")
	var dummy = []string{""}
	attachCmd.Flags().StringSliceVar(&containers, "c", dummy , "Specify container ids")	
	healCmd.Flags().StringSliceVar(&containers, "c", dummy , "Specify container ids")	
	autoHealCmd.Flags().StringSliceVar(&containers, "c", dummy , "Specify container ids")	
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}else {
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigFile(".cobra")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}