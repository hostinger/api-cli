package profiles

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Profiles commands",
}

func init() {
	GroupCmd.AddCommand(ConnectedSendingDomainCmd)
	GroupCmd.AddCommand(DomainDnsStatusCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListPlanFeatureAccessCmd)
	GroupCmd.AddCommand(RemainingPlanLimitsCmd)
}
