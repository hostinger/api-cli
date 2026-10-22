package domains

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "domains",
	Short: "Domains commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ChangeWebsiteCmd)
	GroupCmd.AddCommand(LinkToWebsiteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UnlinkFromWebsiteCmd)
}
