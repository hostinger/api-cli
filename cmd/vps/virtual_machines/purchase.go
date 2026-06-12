package virtual_machines

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

var PurchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase new virtual machine",
	Long:  "Purchase and setup a new virtual machine.\n\nIf virtual machine setup fails for any reason, login to\n[hPanel](https://hpanel.hostinger.com/) and complete the setup manually.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nUse this endpoint to create new VPS instances.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(purchaseBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSPurchaseNewVirtualMachineV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON)")
	PurchaseCmd.Flags().StringP("item-id", "", "", "Catalog price item ID")
	PurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	PurchaseCmd.Flags().StringP("setup", "", "", " (JSON)")
	PurchaseCmd.MarkFlagRequired("item-id")
	PurchaseCmd.MarkFlagRequired("setup")
}

func purchaseBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		v, _ := cmd.Flags().GetString("coupons")
		body["coupons"] = utils.JSONValue(v, "coupons")
	}
	itemIdVal, _ := cmd.Flags().GetString("item-id")
	body["item_id"] = itemIdVal
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	setupVal, _ := cmd.Flags().GetString("setup")
	body["setup"] = utils.JSONValue(setupVal, "setup")
	return body
}
