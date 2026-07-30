package subscriptions

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

var RenewCmd = &cobra.Command{
	Use:   "renew <subscription-id>",
	Short: "Renew subscription",
	Long:  "Create a renewal order for an existing Hostinger subscription.\n\nThis endpoint places a renewal order for a single subscription, leveraging\nthe existing billing infrastructure. Use the\n[subscriptions endpoint](#tag/billing-subscriptions) to look up the\n`subscriptionId` values available for renewal.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nUse this endpoint to renew any subscription available in your account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(renewBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().BillingRenewSubscriptionV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RenewCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON)")
	RenewCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
}

func renewBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		v, _ := cmd.Flags().GetString("coupons")
		body["coupons"] = utils.JSONValue(v, "coupons")
	}
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body
}
