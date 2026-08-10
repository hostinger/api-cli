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

var ClaimFreeCmd = &cobra.Command{
	Use:   "claim-free",
	Short: "Claim free domain",
	Long:  "Claim a free domain available on your account and register it.\n\nUnlike purchasing a domain, this consumes a free domain you already have,\nso no payment method is required.\n\nA successful response means the domain is registered. If registration fails, login to\n[hPanel](https://hpanel.hostinger.com/) and check domain registration status.\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nSome TLDs require `additional_details` to be provided and these will be validated before claiming.\n\nRequests which cannot be fulfilled are rejected with an error code in the response body,\nfor example `2037` when no free domain is available.\n\nUse this endpoint to register a domain using a free domain from your account.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(claimFreeBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsClaimFreeDomainV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ClaimFreeCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON)")
	ClaimFreeCmd.Flags().StringP("domain", "", "", "Domain name")
	ClaimFreeCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON)")
	ClaimFreeCmd.MarkFlagRequired("domain")
}

func claimFreeBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		v, _ := cmd.Flags().GetString("additional-details")
		body["additional_details"] = utils.JSONValue(v, "additional-details")
	}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	if cmd.Flags().Changed("domain-contacts") {
		v, _ := cmd.Flags().GetString("domain-contacts")
		body["domain_contacts"] = utils.JSONValue(v, "domain-contacts")
	}
	return body
}
