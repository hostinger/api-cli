package virtual_machines

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PurchaseCmd = &cobra.Command{
	Use:               "purchase",
	Short:             "Purchase new virtual machine",
	Long:              "Purchase and setup a new virtual machine.\n\nIf virtual machine setup fails for any reason, login to\nhPanel (https://hpanel.hostinger.com/) and complete the setup manually.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nIf the response is `202 Accepted`, the payment is still being processed and the virtual machine\nwas not set up. Login to\nhPanel (https://hpanel.hostinger.com/) and complete the setup manually.\n\nUse this endpoint to create new VPS instances.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --coupons\n    [\n      \"Coupon 3\"\n    ]\n\n  --setup\n    {\n      \"data_center_id\": 19,\n      \"enable_backups\": true,\n      \"hostname\": \"my.server.tld\",\n      \"install_monarx\": false,\n      \"ns1\": \"4.3.2.1\",\n      \"ns2\": \"1.2.3.4\",\n      \"password\": \"oMeNRustosIO\",\n      \"post_install_script_id\": 6324,\n      \"public_key\": {\n        \"key\": \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2X...\",\n        \"name\": \"my-key\"\n      },\n      \"template_id\": 1130\n    }",
	Example:           "  hostinger vps virtual-machines purchase --item-id hostingercom-vps-kvm2-usd-1m --setup @setup.json",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := purchaseBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Purchase new virtual machine", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSPurchaseNewVirtualMachineV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(PurchaseCmd)
	PurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON, @file or -; shape in the help above)")
	PurchaseCmd.Flags().StringP("item-id", "", "", "Catalog price item ID")
	PurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	PurchaseCmd.Flags().StringP("setup", "", "", "(JSON, @file or -; shape in the help above)")
	PurchaseCmd.MarkFlagRequired("item-id")
	PurchaseCmd.MarkFlagRequired("setup")
	PurchaseCmd.RegisterFlagCompletionFunc("coupons", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PurchaseCmd.RegisterFlagCompletionFunc("item-id", cobra.NoFileCompletions)
	PurchaseCmd.RegisterFlagCompletionFunc("payment-method-id", cobra.NoFileCompletions)
	PurchaseCmd.RegisterFlagCompletionFunc("setup", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PurchaseCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func purchaseBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("coupons") {
		couponsVal, _ := cmd.Flags().GetString("coupons")
		couponsValJSON, err := utils.JSONValue(couponsVal, "coupons")
		if err != nil {
			return nil, err
		}
		body["coupons"] = couponsValJSON
	}
	itemIdVal, _ := cmd.Flags().GetString("item-id")
	body["item_id"] = itemIdVal
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	setupVal, _ := cmd.Flags().GetString("setup")
	setupValJSON, err := utils.JSONValue(setupVal, "setup")
	if err != nil {
		return nil, err
	}
	body["setup"] = setupValJSON
	return body, nil
}
