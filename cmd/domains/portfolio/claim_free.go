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

var ClaimFreeCmd = &cobra.Command{
	Use:               "claim-free",
	Short:             "Claim free domain",
	Long:              "Claim a free domain available on your account and register it.\n\nUnlike purchasing a domain, this consumes a free domain you already have,\nso no payment method is required.\n\nA successful response means the domain is registered. If registration fails, login to\nhPanel (https://hpanel.hostinger.com/) and check domain registration status.\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nSome TLDs require `additional_details` to be provided and these will be validated before claiming.\n\nRequests which cannot be fulfilled are rejected with an error code in the response body,\nfor example `2037` when no free domain is available.\n\nUse this endpoint to register a domain using a free domain from your account.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --additional-details\n    {}\n\n  --domain-contacts\n    {\n      \"admin_id\": 546123,\n      \"billing_id\": 741288,\n      \"owner_id\": 741288,\n      \"tech_id\": 741288\n    }",
	Example:           "  hostinger domains portfolio claim-free --domain my-new-domain.tld",
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
			r, err := api.Request().DomainsClaimFreeDomainV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	ClaimFreeCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)")
	ClaimFreeCmd.Flags().StringP("domain", "", "", "Domain name")
	ClaimFreeCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON, @file or -; shape in the help above)")
	ClaimFreeCmd.MarkFlagRequired("domain")
	ClaimFreeCmd.RegisterFlagCompletionFunc("additional-details", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	ClaimFreeCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	ClaimFreeCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func claimFreeBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		additionalDetailsVal, _ := cmd.Flags().GetString("additional-details")
		additionalDetailsValJSON, err := utils.JSONValue(additionalDetailsVal, "additional-details")
		if err != nil {
			return nil, err
		}
		body["additional_details"] = additionalDetailsValJSON
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
	return body, nil
}
