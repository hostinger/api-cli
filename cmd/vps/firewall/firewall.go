package firewall

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Firewall commands",
}

func init() {
	GroupCmd.AddCommand(ActivateCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(CreateRuleCmd)
	GroupCmd.AddCommand(DeactivateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DeleteRuleCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(SyncCmd)
	GroupCmd.AddCommand(UpdateRuleCmd)
}
