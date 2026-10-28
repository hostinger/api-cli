package integrations

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "integrations",
	Short: "Integrations commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ConnectAiWebsiteBuilderSiteCmd)
	GroupCmd.AddCommand(ConnectEcommerceStoreCmd)
	GroupCmd.AddCommand(DisconnectAiWebsiteBuilderSiteCmd)
	GroupCmd.AddCommand(DisconnectEcommerceStoreCmd)
	GroupCmd.AddCommand(DisconnectWordpressSiteCmd)
	GroupCmd.AddCommand(ListConnectedWebsitesCmd)
	GroupCmd.AddCommand(ListEcommerceStoresCmd)
}
