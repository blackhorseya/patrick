package cmd

import (
	"fmt"
	"os"

	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project"
	"github.com/blackhorseya/patrick/internal/pkg/consts"
	"github.com/google/wire"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	projectBiz project.IProjectBiz
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "patrick",
	Short: "One can generate golang project template and function",
	Long: `
██████╗  █████╗ ████████╗██████╗ ██╗ ██████╗██╗  ██╗
██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██║██╔════╝██║ ██╔╝
██████╔╝███████║   ██║   ██████╔╝██║██║     █████╔╝ 
██╔═══╝ ██╔══██║   ██║   ██╔══██╗██║██║     ██╔═██╗ 
██║     ██║  ██║   ██║   ██║  ██║██║╚██████╗██║  ██╗
╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝

One can generate golang project template and function`,
}

// NewRootCmd return *cobra.Command
func NewRootCmd(biz project.IProjectBiz) (*cobra.Command, error) {
	projectBiz = biz

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.yaml)", consts.AppName))

	rootCmd.Version = consts.Version

	return rootCmd, nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + consts.AppName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewRootCmd)
