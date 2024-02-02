package cmd

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/gen-cli/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// 工作目录
var workDir string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gen-cli/config.yaml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		workDir = fmt.Sprintf("%s/.gen-cli", home)

		// Search config in home directory with name ".gen-cli" (without extension).
		viper.AddConfigPath(workDir)
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&config.Config)
	if err != nil {
		fmt.Println("unable to decode into struct", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gen-cli",
	Short: "gen-cli is a very nice generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("hello gen-cli")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
