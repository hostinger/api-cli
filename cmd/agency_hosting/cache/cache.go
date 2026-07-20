package cache

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache commands",
}

func init() {
	GroupCmd.AddCommand(ClearPlanWebsiteCmd)
}
