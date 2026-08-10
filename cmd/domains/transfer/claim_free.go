package transfer

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
	Short: "Claim free domain transfer",
	Long:  "Claim a free domain transfer available on your account and start the transfer.\n\nUnlike purchasing a transfer, this consumes a free domain transfer you already have,\nso no payment method is required.\n\nBefore making request, unlock the domain at the current registrar and get its authorization\ncode. The transfer is validated first, so domains which cannot be transferred are rejected\nbefore the free domain transfer is consumed.\n\nA successful response means the transfer has been started. Completion depends on the current\nregistrar and can be followed with the [transfer list endpoint](#tag/domains-transfer).\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nRequests which cannot be fulfilled are rejected with an error code in the response body.\n\nUse this endpoint to transfer a domain using a free domain transfer from your account.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(claimFreeBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsClaimFreeDomainTransferV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ClaimFreeCmd.Flags().StringP("auth-code", "", "", "Authorization code from the current registrar")
	ClaimFreeCmd.Flags().StringP("domain", "", "", "Domain name")
	ClaimFreeCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON)")
	ClaimFreeCmd.Flags().BoolP("should-keep-ns", "", true, "Keep the existing nameservers of the domain")
	ClaimFreeCmd.MarkFlagRequired("auth-code")
	ClaimFreeCmd.MarkFlagRequired("domain")
}

func claimFreeBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	authCodeVal, _ := cmd.Flags().GetString("auth-code")
	body["auth_code"] = authCodeVal
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	if cmd.Flags().Changed("domain-contacts") {
		v, _ := cmd.Flags().GetString("domain-contacts")
		body["domain_contacts"] = utils.JSONValue(v, "domain-contacts")
	}
	if cmd.Flags().Changed("should-keep-ns") {
		v, _ := cmd.Flags().GetBool("should-keep-ns")
		body["should_keep_ns"] = v
	}
	return body
}
