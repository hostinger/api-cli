package orders

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListStoreCmd = &cobra.Command{
	Use:               "list-store <store_id>",
	Short:             "List store orders",
	Long:              "List a store's orders newest first as summaries. Filter by status, payment or fulfilment\nstatus, customer email, order number or a free-text query. Amounts are in the smallest\ncurrency unit. Retrieve a single order for its line items, addresses and fulfilments.",
	Example:           "  hostinger ecommerce orders list-store store_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listStoreParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceListStoreOrdersV1WithResponse(ctx, args[0], params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(ListStoreCmd, utils.StringArg("store_id"))
	ListStoreCmd.Flags().StringSliceP("status", "", nil, "Order statuses to include. (one of: pending, completed, archived, canceled, requires_action)")
	ListStoreCmd.Flags().StringSliceP("payment-status", "", nil, "Payment statuses to include. A paid order is \"captured\". (one of: not_paid, awaiting, captured, partially_refunded, refunded, canceled, requires_action, not_required)")
	ListStoreCmd.Flags().StringSliceP("fulfillment-status", "", nil, "Fulfilment statuses to include. (one of: not_fulfilled, partially_fulfilled, fulfilled, partially_shipped, shipped, partially_returned, returned, canceled, requires_action)")
	ListStoreCmd.Flags().StringP("email", "", "", "Customer email, matched exactly.")
	ListStoreCmd.Flags().StringP("display-id", "", "", "The order number the merchant and customer see.")
	ListStoreCmd.Flags().StringP("q", "", "", "Free-text search over customer name, email, order number and line items.")
	ListStoreCmd.Flags().StringP("created-at-from", "", "", "Earliest creation time to include, inclusive. Accepts a date or ISO date-time (UTC).")
	ListStoreCmd.Flags().StringP("created-at-to", "", "", "Latest creation time to include, inclusive. A bare date covers that whole day.")
	ListStoreCmd.Flags().IntP("page", "", 0, "Page number")
	ListStoreCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"pending", "completed", "archived", "canceled", "requires_action"}, cobra.ShellCompDirectiveNoFileComp))
	ListStoreCmd.RegisterFlagCompletionFunc("payment-status", cobra.FixedCompletions([]cobra.Completion{"not_paid", "awaiting", "captured", "partially_refunded", "refunded", "canceled", "requires_action", "not_required"}, cobra.ShellCompDirectiveNoFileComp))
	ListStoreCmd.RegisterFlagCompletionFunc("fulfillment-status", cobra.FixedCompletions([]cobra.Completion{"not_fulfilled", "partially_fulfilled", "fulfilled", "partially_shipped", "shipped", "partially_returned", "returned", "canceled", "requires_action"}, cobra.ShellCompDirectiveNoFileComp))
	ListStoreCmd.RegisterFlagCompletionFunc("email", cobra.NoFileCompletions)
	ListStoreCmd.RegisterFlagCompletionFunc("display-id", cobra.NoFileCompletions)
	ListStoreCmd.RegisterFlagCompletionFunc("q", cobra.NoFileCompletions)
	ListStoreCmd.RegisterFlagCompletionFunc("created-at-from", cobra.NoFileCompletions)
	ListStoreCmd.RegisterFlagCompletionFunc("created-at-to", cobra.NoFileCompletions)
	ListStoreCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListStoreCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListStoreCmd.MarkFlagsMutuallyExclusive("all", "page")
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
