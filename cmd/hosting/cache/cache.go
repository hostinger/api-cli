package cache

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache commands",
}

func init() {
	GroupCmd.AddCommand(ClearWebsiteCmd)
	GroupCmd.AddCommand(DisableCachelessCmd)
	GroupCmd.AddCommand(DisableWebsiteCmd)
	GroupCmd.AddCommand(EnableCachelessCmd)
	GroupCmd.AddCommand(EnableWebsiteCmd)
}
