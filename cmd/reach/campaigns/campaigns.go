package campaigns

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "campaigns",
	Short: "Campaigns commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(PerformanceCmd)
}
