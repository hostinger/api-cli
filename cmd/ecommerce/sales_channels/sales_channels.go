package sales_channels

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "sales-channels",
	Short: "Sales channels commands",
}

func init() {
	GroupCmd.AddCommand(CreateCustomCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
