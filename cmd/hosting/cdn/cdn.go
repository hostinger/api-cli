package cdn

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cdn",
	Short: "CDN commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(DisableCmd)
	GroupCmd.AddCommand(EnableCmd)
	GroupCmd.AddCommand(PurgeCacheCmd)
	GroupCmd.AddCommand(StatusCmd)
}
