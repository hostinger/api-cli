package payment_methods

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <payment method ID>",
	Short: "Delete payment method",
	Long:  `This endpoint deletes a payment method. Deleted payment methods can no longer be used for placing new orders.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingDeletePaymentMethodV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
