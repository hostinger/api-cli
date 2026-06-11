package payment_methods

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get payment method list",
	Long:  `This endpoint retrieves a list of available payment methods that can be used for placing new orders.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingGetPaymentMethodListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
