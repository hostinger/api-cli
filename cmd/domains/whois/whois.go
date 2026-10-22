package whois

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "whois",
	Short: "WHOIS commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CancelPendingIrtpVerificationCmd)
	GroupCmd.AddCommand(ChangeForCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(PendingIrtpVerificationCmd)
	GroupCmd.AddCommand(SetAsDefaultCmd)
	GroupCmd.AddCommand(UnsetDefaultCmd)
	GroupCmd.AddCommand(UsageCmd)
}
