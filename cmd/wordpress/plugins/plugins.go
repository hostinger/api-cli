package plugins

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Plugins commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ActivateCmd)
	GroupCmd.AddCommand(CheckIfWooCommerceIsInstalledCmd)
	GroupCmd.AddCommand(DeactivateCmd)
	GroupCmd.AddCommand(DeployCmd)
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListInstalledCmd)
	GroupCmd.AddCommand(ListSuggestedCmd)
	GroupCmd.AddCommand(SearchCmd)
	GroupCmd.AddCommand(UninstallCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(UpdateHostingerCmd)
}
