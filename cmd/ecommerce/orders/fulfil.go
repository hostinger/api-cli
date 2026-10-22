package orders

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var FulfilCmd = &cobra.Command{
	Use:               "fulfil <store_id> <order_id>",
	Short:             "Fulfil an order",
	Long:              "Create a fulfilment for the order and attach tracking in one call. Omit items to fulfil\nevery remaining unfulfilled item. Returns the updated order summary.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --items\n    [\n      {\n        \"line_item_id\": \"item_01J8Z5F8W9K8M4A7B3C2D1E0FG\",\n        \"quantity\": 2\n      }\n    ]",
	Example:           "  hostinger ecommerce orders fulfil store_01J8Z5F8W9K8M4A7B3C2D1E0FG order_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := fulfilBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceFulfilAnOrderV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(FulfilCmd, utils.StringArg("store_id"), utils.StringArg("order_id"))
	FulfilCmd.Flags().StringP("items", "", "", "Line items to fulfil. Omit to fulfil every remaining unfulfilled item. (JSON, @file or -; shape in the help above)")
	FulfilCmd.Flags().BoolP("notify-customer", "", false, "Whether to email the customer about the fulfilment. Defaults to true.")
	FulfilCmd.Flags().StringP("tracking-number", "", "", "Carrier tracking number for the shipment.")
	FulfilCmd.Flags().StringP("tracking-url", "", "", "Public tracking URL for the shipment. Requires tracking_number.")
	FulfilCmd.RegisterFlagCompletionFunc("items", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	FulfilCmd.RegisterFlagCompletionFunc("notify-customer", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	FulfilCmd.RegisterFlagCompletionFunc("tracking-number", cobra.NoFileCompletions)
	FulfilCmd.RegisterFlagCompletionFunc("tracking-url", cobra.NoFileCompletions)
}

func fulfilBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("items") {
		itemsVal, _ := cmd.Flags().GetString("items")
		itemsValJSON, err := utils.JSONValue(itemsVal, "items")
		if err != nil {
			return nil, err
		}
		body["items"] = itemsValJSON
	}
	if cmd.Flags().Changed("notify-customer") {
		v, _ := cmd.Flags().GetBool("notify-customer")
		body["notify_customer"] = v
	}
	if cmd.Flags().Changed("tracking-number") {
		v, _ := cmd.Flags().GetString("tracking-number")
		body["tracking_number"] = v
	}
	if cmd.Flags().Changed("tracking-url") {
		v, _ := cmd.Flags().GetString("tracking-url")
		body["tracking_url"] = v
	}
	return body, nil
}
