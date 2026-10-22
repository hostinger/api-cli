package subscriptions

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RenewCmd = &cobra.Command{
	Use:     "renew <subscription-id>",
	Short:   "Renew subscription",
	Long:    "Create a renewal order for an existing Hostinger subscription.\n\nThis endpoint places a renewal order for a single subscription, leveraging\nthe existing billing infrastructure. Use the\nsubscriptions endpoint (#tag/billing-subscriptions) to look up the\n`subscriptionId` values available for renewal.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nIf the response is `202 Accepted`, the payment is still being processed and the renewal will\ncomplete asynchronously once the payment is confirmed.\n\nUse this endpoint to renew any subscription available in your account.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --coupons\n    [\n      \"Coupon 3\"\n    ]",
	Example: "  hostinger billing subscriptions renew Cxy353Uhl1xC54pG6",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().BillingGetSubscriptionListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := renewBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().BillingRenewSubscriptionV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RenewCmd, utils.StringArg("subscription-id"))
	RenewCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON, @file or -; shape in the help above)")
	RenewCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	RenewCmd.RegisterFlagCompletionFunc("coupons", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	RenewCmd.RegisterFlagCompletionFunc("payment-method-id", cobra.NoFileCompletions)
}

func renewBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		couponsVal, _ := cmd.Flags().GetString("coupons")
		couponsValJSON, err := utils.JSONValue(couponsVal, "coupons")
		if err != nil {
			return nil, err
		}
		body["coupons"] = couponsValJSON
	}
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body, nil
}
