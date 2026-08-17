package metrics

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Metrics commands",
}

func init() {
	GroupCmd.AddCommand(ListOrderResourceUsageCmd)
	GroupCmd.AddCommand(ListPlanOrderDiskUsageCmd)
}
