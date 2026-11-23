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

var UpdateCmd = &cobra.Command{
	Use:               "update <store_id> <endpoint_id>",
	Short:             "Update an ecommerce webhook",
	Long:              "Change the receiver, events, channel filter or enabled state. Omitted fields are preserved; a null\nchannel filter restores store-wide delivery. Endpoints disabled automatically after a 410 Gone\n(receiver_gone) or persistent delivery failures (receiver_failing) resume with enabled: true; events\nfrom the disabled period are not sent. Lifecycle-revoked endpoints cannot be enabled again: delete\nthem and create fresh subscriptions after recovery.\n\n400 codes: webhook_destination_not_allowed (the new URL, or the saved URL when enabling, is not a public\ndestination), webhook_sales_channel_invalid. 409 codes: webhook_endpoint_revoked, webhook_store_inactive,\nwebhook_endpoint_catalog_busy (retry). Errors return the upstream code in `message`, as\n`[Ecommerce:<code>] <reason>`.",
	Example:           "  hostinger ecommerce webhooks update store_01J8Z5F8W9K8M4A7B3C2D1E0FG whep_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceUpdateAnEcommerceWebhookV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("store_id"), utils.StringArg("endpoint_id"))
	UpdateCmd.Flags().BoolP("enabled", "", false, "Merchant-disabled endpoints can be enabled again; lifecycle-revoked endpoints cannot.")
	UpdateCmd.Flags().StringSliceP("events", "", nil, "Replacement event types. Omit to preserve. - `order.placed`: a customer places an order, before any online payment is completed. Orders paid online may not have an order number yet. - `order.paid`: an order becomes fully paid: an online payment is captured, a free order is placed or the merchant marks it as paid. Partial payments don't send it. - `order.cancelled`: an unpaid order is canceled, for example by the merchant, a connected sales channel or after a delayed payment fails. Paid orders can't be canceled, so a cancellation never involves a refund. - `order.refunded`: each time a full or partial refund is recorded, through the payment provider or manually. - `order.fulfilled`: a fulfillment leaves every item in the order fulfilled. Partial fulfillments send `fulfillment.created` instead. - `fulfillment.created`: each time items in an order are fulfilled, including partial fulfillments and digital items fulfilled automatically after payment. - `fulfillment.updated`: a fulfillment's tracking details change, or a connected shipping service or sales channel reports progress such as delivery. - `fulfillment.cancelled`: a fulfillment is canceled before it ships, for example when it is canceled or its shipping label is voided. - `product.created`: a product is added to the store, including drafts and products created in bulk. - `product.updated`: a product changes, including its details, variants, prices, stock, visibility or collections. Stock changes from orders also send it, so expect frequent events. - `product.deleted`: a product is deleted, individually or in bulk. (one of: order.placed, order.paid, order.cancelled, order.refunded, order.fulfilled, fulfillment.created, fulfillment.updated, fulfillment.cancelled, product.created, product.updated, product.deleted)")
	UpdateCmd.Flags().StringSliceP("sales-channel-ids", "", nil, "Omit to preserve; null restores store-wide delivery. Empty or foreign-store IDs are invalid.")
	UpdateCmd.Flags().StringP("url", "", "", "HTTPS receiver on port 443, without credentials or a fragment. Omit to preserve. A changed URL must be a public destination under the same rules as creation; resending the saved URL unchanged is accepted.")
	UpdateCmd.RegisterFlagCompletionFunc("enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"order.placed", "order.paid", "order.cancelled", "order.refunded", "order.fulfilled", "fulfillment.created", "fulfillment.updated", "fulfillment.cancelled", "product.created", "product.updated", "product.deleted"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("sales-channel-ids", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("enabled") {
		v, _ := cmd.Flags().GetBool("enabled")
		body["enabled"] = v
	}
	if cmd.Flags().Changed("events") {
		v, _ := cmd.Flags().GetStringSlice("events")
		body["events"] = v
	}
	if cmd.Flags().Changed("sales-channel-ids") {
		v, _ := cmd.Flags().GetStringSlice("sales-channel-ids")
		body["sales_channel_ids"] = v
	}
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body, nil
}
