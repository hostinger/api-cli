package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreatePurchaseCmd = &cobra.Command{
	Use:   "create-purchase",
	Short: "Create purchase order",
	Long:  "Create a purchase order for any Hostinger product.\n\nThis unified endpoint places an order for one or more catalog items and\nworks across all Hostinger products, leveraging the existing billing\ninfrastructure. Use the [catalog endpoint](#tag/billing-catalog) to look\nup the `item_id` values available for purchase.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nThis endpoint only places the order. Product-specific provisioning\n(e.g. VPS setup or domain registration) is not performed here — once the\norder completes, use the relevant product endpoints or\n[hPanel](https://hpanel.hostinger.com/) to finalize setup.\n\nUse this endpoint to purchase any product available in the catalog.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createPurchaseBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().BillingCreatePurchaseOrderV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreatePurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON)")
	CreatePurchaseCmd.Flags().StringP("items", "", "", "Catalog price items to purchase (JSON)")
	CreatePurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	CreatePurchaseCmd.MarkFlagRequired("items")
}

func createPurchaseBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		v, _ := cmd.Flags().GetString("coupons")
		body["coupons"] = utils.JSONValue(v, "coupons")
	}
	itemsVal, _ := cmd.Flags().GetString("items")
	body["items"] = utils.JSONValue(itemsVal, "items")
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body
}
