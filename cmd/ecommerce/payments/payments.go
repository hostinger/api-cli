package payments

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "payments",
	Short: "Payments commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateProviderConnectLinkCmd)
	GroupCmd.AddCommand(EnableManualMethodCmd)
	GroupCmd.AddCommand(ListStoreProvidersCmd)
}
