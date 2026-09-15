package git

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "git",
	Short: "Git commands",
}

func init() {
	GroupCmd.AddCommand(ListInstallationRepositoriesCmd)
	GroupCmd.AddCommand(ListInstallationsCmd)
}
