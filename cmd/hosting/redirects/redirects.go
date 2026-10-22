package redirects

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "redirects",
	Short: "Redirects commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
