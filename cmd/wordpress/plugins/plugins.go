package plugins

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Plugins commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
}
