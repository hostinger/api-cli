package portfolio

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
	Short:             "Purchase new domain",
	Long:              "Purchase and register a new domain name.\n\nIf registration fails, login to hPanel (https://hpanel.hostinger.com/) and check domain registration status.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nIf the response is `202 Accepted`, the payment is still being processed and the domain was\nnot registered. Once the order completes, register the domain from\nhPanel (https://hpanel.hostinger.com/).\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nSome TLDs require `additional_details` to be provided and these will be validated before completing purchase.\n\nUse this endpoint to register new domains for users.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --additional-details\n    {}\n\n  --coupons\n    [\n      \"Coupon 3\"\n    ]\n\n  --domain-contacts\n    {\n      \"admin_id\": 546123,\n      \"billing_id\": 741288,\n      \"owner_id\": 741288,\n      \"tech_id\": 741288\n    }",
	Example:           "  hostinger domains portfolio purchase --domain my-new-domain.tld --item-id hostingercom-domain-com-usd-1y",
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
		if err := utils.Confirm(cmd, "Purchase new domain", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsPurchaseNewDomainV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	PurchaseCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)")
	PurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON, @file or -; shape in the help above)")
	PurchaseCmd.Flags().StringP("domain", "", "", "Domain name")
	PurchaseCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON, @file or -; shape in the help above)")
	PurchaseCmd.Flags().StringP("item-id", "", "", "Catalog price item ID")
	PurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	PurchaseCmd.MarkFlagRequired("domain")
	PurchaseCmd.MarkFlagRequired("item-id")
	PurchaseCmd.RegisterFlagCompletionFunc("additional-details", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PurchaseCmd.RegisterFlagCompletionFunc("coupons", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PurchaseCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	PurchaseCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PurchaseCmd.RegisterFlagCompletionFunc("item-id", cobra.NoFileCompletions)
	PurchaseCmd.RegisterFlagCompletionFunc("payment-method-id", cobra.NoFileCompletions)
	PurchaseCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func purchaseBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		additionalDetailsVal, _ := cmd.Flags().GetString("additional-details")
		additionalDetailsValJSON, err := utils.JSONValue(additionalDetailsVal, "additional-details")
		if err != nil {
			return nil, err
		}
		body["additional_details"] = additionalDetailsValJSON
	}
	if cmd.Flags().Changed("coupons") {
		couponsVal, _ := cmd.Flags().GetString("coupons")
		couponsValJSON, err := utils.JSONValue(couponsVal, "coupons")
		if err != nil {
			return nil, err
		}
		body["coupons"] = couponsValJSON
	}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	if cmd.Flags().Changed("domain-contacts") {
		domainContactsVal, _ := cmd.Flags().GetString("domain-contacts")
		domainContactsValJSON, err := utils.JSONValue(domainContactsVal, "domain-contacts")
		if err != nil {
			return nil, err
		}
		body["domain_contacts"] = domainContactsValJSON
	}
	itemIdVal, _ := cmd.Flags().GetString("item-id")
	body["item_id"] = itemIdVal
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body, nil
}
