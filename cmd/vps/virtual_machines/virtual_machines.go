package virtual_machines

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:     "virtual-machines",
	Aliases: []string{"vm"},
	Short:   "Virtual machine commands",
}

func init() {
	GroupCmd.AddCommand(AttachedPublicKeysCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(MetricsCmd)
	GroupCmd.AddCommand(PurchaseCmd)
	GroupCmd.AddCommand(RecreateCmd)
	GroupCmd.AddCommand(ResetHostnameCmd)
	GroupCmd.AddCommand(RestartCmd)
	GroupCmd.AddCommand(SetHostnameCmd)
	GroupCmd.AddCommand(SetNameserversCmd)
	GroupCmd.AddCommand(SetPanelPasswordCmd)
	GroupCmd.AddCommand(SetRootPasswordCmd)
	GroupCmd.AddCommand(SetupCmd)
	GroupCmd.AddCommand(StartCmd)
	GroupCmd.AddCommand(StopCmd)
}
