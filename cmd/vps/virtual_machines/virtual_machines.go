package virtual_machines

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "vm",
	Short: "Virtual machine management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(GetAttachedKeysCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(GetMetricsCmd)
	GroupCmd.AddCommand(PurchaseCmd)
	GroupCmd.AddCommand(SetupCmd)
	GroupCmd.AddCommand(StartCmd)
	GroupCmd.AddCommand(StopCmd)
	GroupCmd.AddCommand(RestartCmd)
	GroupCmd.AddCommand(SetHostnameCmd)
	GroupCmd.AddCommand(SetNameserversCmd)
	GroupCmd.AddCommand(SetPanelPasswordCmd)
	GroupCmd.AddCommand(SetRootPasswordCmd)
	GroupCmd.AddCommand(ResetHostnameCmd)
	GroupCmd.AddCommand(RecreateCmd)
}
