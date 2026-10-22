package ssl

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ssl",
	Short: "SSL commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(StatusCmd)
	GroupCmd.AddCommand(ToggleHttpsRedirectCmd)
	GroupCmd.AddCommand(UninstallCmd)
}
