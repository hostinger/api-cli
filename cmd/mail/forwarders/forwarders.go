package forwarders

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "forwarders",
	Short: "Forwarders commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ResendConfirmationCmd)
	GroupCmd.AddCommand(UpdateKeepCopySettingCmd)
}
