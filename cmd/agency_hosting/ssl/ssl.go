package ssl

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ssl",
	Short: "SSL commands",
}

func init() {
	GroupCmd.AddCommand(InstallWebsiteCmd)
	GroupCmd.AddCommand(ReinstallWebsiteCmd)
	GroupCmd.AddCommand(UninstallWebsiteCmd)
	GroupCmd.AddCommand(WebsiteStatusCmd)
}
