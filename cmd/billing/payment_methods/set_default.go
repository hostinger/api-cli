package payment_methods

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var SetDefaultCmd = &cobra.Command{
	Use:   "set-default <payment method ID>",
	Short: "Set default payment method",
	Long:  `This endpoint sets a specified payment method as the default. The default payment method is used automatically for new orders.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingSetDefaultPaymentMethodV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
