package payment_methods

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetDefaultCmd = &cobra.Command{
	Use:   "set-default <payment-method-id>",
	Short: "Set default payment method",
	Long:  "Set the default payment method for your account.\n\nUse this endpoint to configure the primary payment method for future orders.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingSetDefaultPaymentMethodV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
