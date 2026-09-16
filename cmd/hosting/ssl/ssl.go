package ssl

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ssl",
	Short: "SSL commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(StatusCmd)
	GroupCmd.AddCommand(ToggleHttpsRedirectCmd)
	GroupCmd.AddCommand(UninstallCmd)
}
