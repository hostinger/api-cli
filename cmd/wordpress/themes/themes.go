package themes

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "themes",
	Short: "Themes commands",
}

func init() {
	GroupCmd.AddCommand(ActivateCmd)
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListInstalledCmd)
	GroupCmd.AddCommand(UninstallCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
