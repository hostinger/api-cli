package payment_methods

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get payment method list",
	Long:  "Retrieve available payment methods that can be used for placing new orders.\n\nIf you want to add new payment method,\nplease use [hPanel](https://hpanel.hostinger.com/billing/payment-methods).\n\nUse this endpoint to view available payment options before creating orders.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingGetPaymentMethodListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
