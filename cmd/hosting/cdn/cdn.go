package cdn

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cdn",
	Short: "CDN commands",
}

func init() {
	GroupCmd.AddCommand(DisableCmd)
	GroupCmd.AddCommand(EnableCmd)
	GroupCmd.AddCommand(PurgeCacheCmd)
	GroupCmd.AddCommand(StatusCmd)
}
