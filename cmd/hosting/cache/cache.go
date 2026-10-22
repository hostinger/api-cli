package cache

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ClearWebsiteCmd)
	GroupCmd.AddCommand(ToggleCachelessCmd)
	GroupCmd.AddCommand(ToggleWebsiteCmd)
}
