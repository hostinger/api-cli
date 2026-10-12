package profiles

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Profiles commands",
}

func init() {
	GroupCmd.AddCommand(ConnectSendingDomainCmd)
	GroupCmd.AddCommand(ConnectedSendingDomainCmd)
	GroupCmd.AddCommand(DeleteSendingDomainCmd)
	GroupCmd.AddCommand(DomainDnsStatusCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListPlanFeatureAccessCmd)
	GroupCmd.AddCommand(RemainingPlanLimitsCmd)
	GroupCmd.AddCommand(SettingsCmd)
	GroupCmd.AddCommand(UpdateSettingsCmd)
}
