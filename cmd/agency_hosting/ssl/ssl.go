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
	GroupCmd.AddCommand(InstallWebsiteCmd)
	GroupCmd.AddCommand(ReinstallWebsiteCmd)
	GroupCmd.AddCommand(UninstallWebsiteCmd)
	GroupCmd.AddCommand(WebsiteStatusCmd)
}
