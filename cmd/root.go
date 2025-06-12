package cmd

import (
	"fmt"
	"github.com/hostinger/api-cli/cmd/vps"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	OutputFormat string
	validFormats = []string{"json", "table", "tree"}
)

var RootCmd = &cobra.Command{
	Use:   "hapi",
	Short: "Hostinger API Command Line Interface",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.DisableAutoGenTag = true

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.hapi.yaml)")
	RootCmd.PersistentFlags().StringVar(&OutputFormat, "format", "", "Output format type (json|table|tree), default: table")

	RootCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return validFormats, cobra.ShellCompDirectiveNoFileComp
	})

	RootCmd.AddCommand(vps.GroupCmd)
	RootCmd.AddCommand(VersionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hapi")
	}

	viper.SetEnvPrefix("hapi")
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
