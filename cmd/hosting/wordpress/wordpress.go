package wordpress

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "Wordpress commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ListInstallationsCmd)
}
