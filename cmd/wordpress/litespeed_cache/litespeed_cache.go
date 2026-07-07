package litespeed_cache

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "litespeed-cache",
	Short: "LiteSpeed Cache commands",
}

func init() {
	GroupCmd.AddCommand(PurgeLiteSpeedCmd)
	GroupCmd.AddCommand(ShowLiteSpeedStatusCmd)
}
