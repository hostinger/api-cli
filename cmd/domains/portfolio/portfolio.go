package portfolio

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Portfolio commands",
}

func init() {
	GroupCmd.AddCommand(DisableLockCmd)
	GroupCmd.AddCommand(DisablePrivacyProtectionCmd)
	GroupCmd.AddCommand(EnableLockCmd)
	GroupCmd.AddCommand(EnablePrivacyProtectionCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(PurchaseCmd)
	GroupCmd.AddCommand(UpdateNameserversCmd)
}
