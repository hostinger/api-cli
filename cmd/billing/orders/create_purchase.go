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

var CreatePurchaseCmd = &cobra.Command{
	Use:               "create-purchase",
	Short:             "Create purchase order",
	Long:              "Create a purchase order for any Hostinger product.\n\nThis unified endpoint places an order for one or more catalog items and\nworks across all Hostinger products, leveraging the existing billing\ninfrastructure. Use the catalog endpoint (#tag/billing-catalog) to look\nup the `item_id` values available for purchase.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nIf the response is `202 Accepted`, the payment is still being processed and the order will\ncomplete asynchronously once the payment is confirmed.\n\nThis endpoint only places the order. Product-specific provisioning\n(e.g. VPS setup or domain registration) is not performed here — once the\norder completes, use the relevant product endpoints or\nhPanel (https://hpanel.hostinger.com/) to finalize setup.\n\nUse this endpoint to purchase any product available in the catalog.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --coupons\n    [\n      \"Coupon 3\"\n    ]\n\n  --items\n    [\n      {\n        \"item_id\": \"hostingercom-vps-kvm2-usd-1m\",\n        \"quantity\": 1\n      }\n    ]",
	Example:           "  hostinger billing orders create-purchase --items '[{\"item_id\":\"hostingercom-vps-kvm2-usd-1m\",\"quantity\":1}]'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createPurchaseBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().BillingCreatePurchaseOrderV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreatePurchaseCmd)
	CreatePurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON, @file or -; shape in the help above)")
	CreatePurchaseCmd.Flags().StringP("items", "", "", "Catalog price items to purchase (JSON, @file or -; shape in the help above)")
	CreatePurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	CreatePurchaseCmd.MarkFlagRequired("items")
	CreatePurchaseCmd.RegisterFlagCompletionFunc("coupons", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreatePurchaseCmd.RegisterFlagCompletionFunc("items", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreatePurchaseCmd.RegisterFlagCompletionFunc("payment-method-id", cobra.NoFileCompletions)
}

func createPurchaseBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		couponsVal, _ := cmd.Flags().GetString("coupons")
		couponsValJSON, err := utils.JSONValue(couponsVal, "coupons")
		if err != nil {
			return nil, err
		}
		body["coupons"] = couponsValJSON
	}
	itemsVal, _ := cmd.Flags().GetString("items")
	itemsValJSON, err := utils.JSONValue(itemsVal, "items")
	if err != nil {
		return nil, err
	}
	body["items"] = itemsValJSON
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body, nil
}
