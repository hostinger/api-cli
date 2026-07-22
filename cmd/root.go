package cmd

import (
	"fmt"
	"github.com/hostinger/api-cli/cmd/agency_hosting"
	"github.com/hostinger/api-cli/cmd/billing"
	"github.com/hostinger/api-cli/cmd/dns"
	"github.com/hostinger/api-cli/cmd/domains"
	"github.com/hostinger/api-cli/cmd/ecommerce"
	"github.com/hostinger/api-cli/cmd/horizons"
	"github.com/hostinger/api-cli/cmd/hosting"
	"github.com/hostinger/api-cli/cmd/mail"
	"github.com/hostinger/api-cli/cmd/reach"
	"github.com/hostinger/api-cli/cmd/vps"
	"github.com/hostinger/api-cli/cmd/wordpress"
	"os"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	OutputFormat string
	validFormats = []string{"json", "table", "tree"}
)

var RootCmd = &cobra.Command{
	Use:   "hostinger",
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.hostinger.yaml)")
	RootCmd.PersistentFlags().StringVar(&OutputFormat, "format", "", "Output format type (json|table|tree), default: table")

	RootCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return validFormats, cobra.ShellCompDirectiveNoFileComp
	})

	RootCmd.AddCommand(agency_hosting.GroupCmd)
	RootCmd.AddCommand(billing.GroupCmd)
	RootCmd.AddCommand(dns.GroupCmd)
	RootCmd.AddCommand(domains.GroupCmd)
	RootCmd.AddCommand(ecommerce.GroupCmd)
	RootCmd.AddCommand(horizons.GroupCmd)
	RootCmd.AddCommand(hosting.GroupCmd)
	RootCmd.AddCommand(mail.GroupCmd)
	RootCmd.AddCommand(reach.GroupCmd)
	RootCmd.AddCommand(vps.GroupCmd)
	RootCmd.AddCommand(wordpress.GroupCmd)
	RootCmd.AddCommand(VersionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		path, _ := utils.MigrateLegacyConfig(home, os.Stderr)
		viper.SetConfigFile(path)
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("hostinger")
	viper.AutomaticEnv() // read in environment variables that match
	utils.BindLegacyEnv(os.Stderr)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
