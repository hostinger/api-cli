package orders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RetrieveCmd = &cobra.Command{
	Use:   "retrieve <store_id> <order_id>",
	Short: "Retrieve an order",
	Long:  "Retrieve one order in full: line items (each with the id the fulfil endpoint needs),\naddresses, the totals breakdown and fulfilments with tracking. Amounts are in the\nsmallest currency unit.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceRetrieveAnOrderV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
