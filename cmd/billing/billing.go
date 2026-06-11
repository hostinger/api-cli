package billing

import (
	"github.com/hostinger/api-cli/cmd/billing/catalog"
	"github.com/hostinger/api-cli/cmd/billing/payment_methods"
	"github.com/hostinger/api-cli/cmd/billing/subscriptions"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "billing",
	Short: "Billing management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(catalog.GroupCmd)
	GroupCmd.AddCommand(payment_methods.GroupCmd)
	GroupCmd.AddCommand(subscriptions.GroupCmd)
}
