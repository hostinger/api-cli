package orders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListStoreCmd = &cobra.Command{
	Use:   "list-store <store_id>",
	Short: "List store orders",
	Long:  "List a store's orders newest first as summaries. Filter by status, payment or fulfilment\nstatus, customer email, order number or a free-text query. Amounts are in the smallest\ncurrency unit. Retrieve a single order for its line items, addresses and fulfilments.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceListStoreOrdersV1WithResponse(context.TODO(), args[0], listStoreParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListStoreCmd.Flags().StringSliceP("status", "", nil, "Order statuses to include. (one of: pending, completed, archived, canceled, requires_action)")
	ListStoreCmd.Flags().StringSliceP("payment-status", "", nil, "Payment statuses to include. A paid order is \"captured\". (one of: not_paid, awaiting, captured, partially_refunded, refunded, canceled, requires_action, not_required)")
	ListStoreCmd.Flags().StringSliceP("fulfillment-status", "", nil, "Fulfilment statuses to include. (one of: not_fulfilled, partially_fulfilled, fulfilled, partially_shipped, shipped, partially_returned, returned, canceled, requires_action)")
	ListStoreCmd.Flags().StringP("email", "", "", "Customer email, matched exactly.")
	ListStoreCmd.Flags().StringP("display-id", "", "", "The order number the merchant and customer see.")
	ListStoreCmd.Flags().StringP("q", "", "", "Free-text search over customer name, email, order number and line items.")
	ListStoreCmd.Flags().StringP("created-at-from", "", "", "Earliest creation time to include, inclusive. Accepts a date or ISO date-time (UTC).")
	ListStoreCmd.Flags().StringP("created-at-to", "", "", "Latest creation time to include, inclusive. A bare date covers that whole day.")
	ListStoreCmd.Flags().IntP("page", "", 0, "Page number")
}

func listStoreParams(cmd *cobra.Command) *client.EcommerceListStoreOrdersV1Params {
	params := &client.EcommerceListStoreOrdersV1Params{}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetStringSlice("status")
		es := make([]client.EcommerceListStoreOrdersV1ParamsStatus, len(v))
		for i, s := range v {
			es[i] = client.EcommerceListStoreOrdersV1ParamsStatus(s)
		}
		params.Status = &es
	}
	if cmd.Flags().Changed("payment-status") {
		v, _ := cmd.Flags().GetStringSlice("payment-status")
		es := make([]client.EcommerceListStoreOrdersV1ParamsPaymentStatus, len(v))
		for i, s := range v {
			es[i] = client.EcommerceListStoreOrdersV1ParamsPaymentStatus(s)
		}
		params.PaymentStatus = &es
	}
	if cmd.Flags().Changed("fulfillment-status") {
		v, _ := cmd.Flags().GetStringSlice("fulfillment-status")
		es := make([]client.EcommerceListStoreOrdersV1ParamsFulfillmentStatus, len(v))
		for i, s := range v {
			es[i] = client.EcommerceListStoreOrdersV1ParamsFulfillmentStatus(s)
		}
		params.FulfillmentStatus = &es
	}
	if cmd.Flags().Changed("email") {
		v, _ := cmd.Flags().GetString("email")
		params.Email = &v
	}
	if cmd.Flags().Changed("display-id") {
		v, _ := cmd.Flags().GetString("display-id")
		params.DisplayId = &v
	}
	if cmd.Flags().Changed("q") {
		v, _ := cmd.Flags().GetString("q")
		params.Q = &v
	}
	if cmd.Flags().Changed("created-at-from") {
		v, _ := cmd.Flags().GetString("created-at-from")
		params.CreatedAtFrom = &v
	}
	if cmd.Flags().Changed("created-at-to") {
		v, _ := cmd.Flags().GetString("created-at-to")
		params.CreatedAtTo = &v
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
