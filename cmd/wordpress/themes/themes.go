package themes

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "themes",
	Short: "Themes commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ActivateCmd)
	GroupCmd.AddCommand(DeployCmd)
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListInstalledCmd)
	GroupCmd.AddCommand(UninstallCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
