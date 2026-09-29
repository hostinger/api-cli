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

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start domain transfer",
	Long:  "Transfer a domain from another registrar to your account.\n\nThe transfer runs on a domain transfer service you have already purchased.\n\nBefore making request, unlock the domain at the current registrar and get its authorization\ncode.\n\nA successful response means the transfer has been started. Completion depends on the current\nregistrar and can be followed with the [transfer list endpoint](#tag/domains-transfer).\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nUse this endpoint to bring domains registered elsewhere into your account.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(startBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsStartDomainTransferV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	StartCmd.Flags().StringP("auth-code", "", "", "Authorization code from the current registrar")
	StartCmd.Flags().StringP("domain", "", "", "Domain name")
	StartCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON)")
	StartCmd.Flags().BoolP("should-keep-ns", "", true, "Keep the existing nameservers of the domain")
	StartCmd.MarkFlagRequired("auth-code")
	StartCmd.MarkFlagRequired("domain")
}

func startBody(cmd *cobra.Command) map[string]any {
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
