package sales_channels

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "sales-channels",
	Short: "Sales channels commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
