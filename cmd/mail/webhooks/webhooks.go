package webhooks

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "webhooks",
	Short: "Webhooks commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListDeliveryLogsCmd)
	GroupCmd.AddCommand(RegenerateSecretCmd)
	GroupCmd.AddCommand(TestCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
