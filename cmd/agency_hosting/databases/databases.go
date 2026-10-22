package databases

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "databases",
	Short: "Databases commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(CreateWebsiteUserCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteUserCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
