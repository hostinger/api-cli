package portfolio

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
	Short: "Purchase new domain",
	Long:  "Purchase and register a new domain name.\n\nIf registration fails, login to [hPanel](https://hpanel.hostinger.com/) and check domain registration status.\n\nIf no payment method is provided, your default payment method will be used automatically.\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nSome TLDs require `additional_details` to be provided and these will be validated before completing purchase.\n\nUse this endpoint to register new domains for users.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(purchaseBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsPurchaseNewDomainV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PurchaseCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON)")
	PurchaseCmd.Flags().StringP("coupons", "", "", "Discount coupon codes (JSON)")
	PurchaseCmd.Flags().StringP("domain", "", "", "Domain name")
	PurchaseCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON)")
	PurchaseCmd.Flags().StringP("item-id", "", "", "Catalog price item ID")
	PurchaseCmd.Flags().IntP("payment-method-id", "", 0, "Payment method ID, default will be used if not provided")
	PurchaseCmd.MarkFlagRequired("domain")
	PurchaseCmd.MarkFlagRequired("item-id")
}

func purchaseBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		v, _ := cmd.Flags().GetString("additional-details")
		body["additional_details"] = utils.JSONValue(v, "additional-details")
	}
	if cmd.Flags().Changed("coupons") {
		v, _ := cmd.Flags().GetString("coupons")
		body["coupons"] = utils.JSONValue(v, "coupons")
	}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	if cmd.Flags().Changed("domain-contacts") {
		v, _ := cmd.Flags().GetString("domain-contacts")
		body["domain_contacts"] = utils.JSONValue(v, "domain-contacts")
	}
	itemIdVal, _ := cmd.Flags().GetString("item-id")
	body["item_id"] = itemIdVal
	if cmd.Flags().Changed("payment-method-id") {
		v, _ := cmd.Flags().GetInt("payment-method-id")
		body["payment_method_id"] = v
	}
	return body
}
