package litespeed_cache

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "litespeed-cache",
	Short: "LiteSpeed Cache commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(PurgeLiteSpeedCmd)
	GroupCmd.AddCommand(ShowLiteSpeedStatusCmd)
}
