package webhooks

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create <store_id>",
	Short:             "Create an ecommerce webhook",
	Long:              "Create an enabled subscription for your store to receive ecommerce events in external automation\nsystems. Save the whsec_ signing secret securely; it is returned once. Up to ten endpoints, including\ndisabled endpoints, may exist per store. Deliveries follow Standard Webhooks: verify the webhook-id,\nwebhook-timestamp and webhook-signature headers against the raw body with an official Standard Webhooks\nlibrary before parsing. Events can arrive out of order or more than once; deduplicate by event id.\n\n400: the URL is not a public destination, or webhook_sales_channel_invalid (a channel ID does not\nbelong to this store). 409 codes: webhook_endpoint_limit_reached, webhook_store_inactive,\nwebhook_endpoint_catalog_busy (a catalog change is in progress; retry). Errors return the upstream code in\n`message`, as `[Ecommerce:<code>] <reason>`.",
	Example:           "  hostinger ecommerce webhooks create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --events order.placed,product.updated --url https://example.com/webhooks/hostinger",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateAnEcommerceWebhookV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(CreateCmd, utils.StringArg("store_id"))
	CreateCmd.Flags().StringSliceP("events", "", nil, "Event types to subscribe to. - `order.placed`: a customer places an order, before any online payment is completed. Orders paid online may not have an order number yet. - `order.paid`: an order becomes fully paid: an online payment is captured, a free order is placed or the merchant marks it as paid. Partial payments don't send it. - `order.cancelled`: an unpaid order is canceled, for example by the merchant, a connected sales channel or after a delayed payment fails. Paid orders can't be canceled, so a cancellation never involves a refund. - `order.refunded`: each time a full or partial refund is recorded, through the payment provider or manually. - `order.fulfilled`: a fulfillment leaves every item in the order fulfilled. Partial fulfillments send `fulfillment.created` instead. - `fulfillment.created`: each time items in an order are fulfilled, including partial fulfillments and digital items fulfilled automatically after payment. - `fulfillment.updated`: a fulfillment's tracking details change, or a connected shipping service or sales channel reports progress such as delivery. - `fulfillment.cancelled`: a fulfillment is canceled before it ships, for example when it is canceled or its shipping label is voided. - `product.created`: a product is added to the store, including drafts and products created in bulk. - `product.updated`: a product changes, including its details, variants, prices, stock, visibility or collections. Stock changes from orders also send it, so expect frequent events. - `product.deleted`: a product is deleted, individually or in bulk. (one of: order.placed, order.paid, order.cancelled, order.refunded, order.fulfilled, fulfillment.created, fulfillment.updated, fulfillment.cancelled, product.created, product.updated, product.deleted)")
	CreateCmd.Flags().StringSliceP("sales-channel-ids", "", nil, "Omit for store-wide delivery. IDs must belong to this store; empty arrays are invalid.")
	CreateCmd.Flags().StringP("url", "", "", "HTTPS receiver on port 443, without credentials or a fragment. Must be a public destination: localhost, single-label and local-network names (.local, .home.arpa, .internal), Hostinger service domains and loopback, private, link-local or reserved IP addresses are rejected. Merchant-owned domains hosted by Hostinger are allowed.")
	CreateCmd.MarkFlagRequired("events")
	CreateCmd.MarkFlagRequired("url")
	CreateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"order.placed", "order.paid", "order.cancelled", "order.refunded", "order.fulfilled", "fulfillment.created", "fulfillment.updated", "fulfillment.cancelled", "product.created", "product.updated", "product.deleted"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("sales-channel-ids", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	eventsVal, _ := cmd.Flags().GetStringSlice("events")
	body["events"] = eventsVal
	if cmd.Flags().Changed("sales-channel-ids") {
		v, _ := cmd.Flags().GetStringSlice("sales-channel-ids")
		body["sales_channel_ids"] = v
	}
	urlVal, _ := cmd.Flags().GetString("url")
	body["url"] = urlVal
	return body, nil
}
