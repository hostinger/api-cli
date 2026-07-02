package themes

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "themes",
	Short: "Themes commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
}
