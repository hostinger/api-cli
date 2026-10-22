package transfer

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ClaimFreeCmd = &cobra.Command{
	Use:               "claim-free",
	Short:             "Claim free domain transfer",
	Long:              "Claim a free domain transfer available on your account and start the transfer.\n\nUnlike purchasing a transfer, this consumes a free domain transfer you already have,\nso no payment method is required.\n\nBefore making request, unlock the domain at the current registrar and get its authorization\ncode. The transfer is validated first, so domains which cannot be transferred are rejected\nbefore the free domain transfer is consumed.\n\nA successful response means the transfer has been started. Completion depends on the current\nregistrar and can be followed with the transfer list endpoint (#tag/domains-transfer).\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nRequests which cannot be fulfilled are rejected with an error code in the response body.\n\nUse this endpoint to transfer a domain using a free domain transfer from your account.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --domain-contacts\n    {\n      \"admin_id\": 546123,\n      \"billing_id\": 741288,\n      \"owner_id\": 741288,\n      \"tech_id\": 741288\n    }",
	Example:           "  hostinger domains transfer claim-free --auth-code Auth123Code456 --domain my-new-domain.tld",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := claimFreeBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsClaimFreeDomainTransferV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ClaimFreeCmd)
	ClaimFreeCmd.Flags().StringP("auth-code", "", "", "Authorization code from the current registrar")
	ClaimFreeCmd.Flags().StringP("domain", "", "", "Domain name")
	ClaimFreeCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON, @file or -; shape in the help above)")
	ClaimFreeCmd.Flags().BoolP("should-keep-ns", "", true, "Keep the existing nameservers of the domain")
	ClaimFreeCmd.MarkFlagRequired("auth-code")
	ClaimFreeCmd.MarkFlagRequired("domain")
	ClaimFreeCmd.RegisterFlagCompletionFunc("auth-code", cobra.NoFileCompletions)
	ClaimFreeCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	ClaimFreeCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	ClaimFreeCmd.RegisterFlagCompletionFunc("should-keep-ns", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func claimFreeBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	authCodeVal, _ := cmd.Flags().GetString("auth-code")
	body["auth_code"] = authCodeVal
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
	if cmd.Flags().Changed("should-keep-ns") {
		v, _ := cmd.Flags().GetBool("should-keep-ns")
		body["should_keep_ns"] = v
	}
	return body, nil
}
