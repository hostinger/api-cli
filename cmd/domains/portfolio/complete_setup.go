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

var CompleteSetupCmd = &cobra.Command{
	Use:     "complete-setup <domain>",
	Short:   "Complete domain setup",
	Long:    "Register a domain you have already paid for but which has not been set up yet.\n\nUse this endpoint when an order completed without registering the domain, for example when\n`Purchase new domain` returned `202 Accepted` and the domain was added to your account without\nbeing registered, or when an earlier setup attempt failed. No new order is placed and no payment\nis taken: the subscription you already own is used, for the period you already paid for.\n\nA domain is left awaiting setup when the details needed to register it were missing or invalid\nas the order completed. Domains ordered elsewhere can be awaiting setup for the same reason.\nComplete the missing information, then call this endpoint. If the order itself has not completed\nyet, the domain is not on your account, wait until it appears in `Get domain list`.\n\nIf `domain_contacts` is omitted, the default WHOIS profile of that TLD is used for all four\nroles. The profile must exist and be complete for the TLD, an incomplete profile is the most\ncommon reason a domain is left awaiting setup. Create one with `Create WHOIS profile`.\n\nSome TLDs require `additional_details`. These are validated before setup, so a missing or\ninvalid value is rejected without any registration being attempted.\n\nThe domain is set up with the default nameservers and without privacy protection. Use\n`Update domain nameservers` and `Enable privacy protection` afterwards to change either.\n\nA successful response means the setup request was accepted, not that the domain is already\nregistered. Poll `Get domain list` for the outcome, the domain appears in `Get domain details`\nonly once it is registered.\n\nUse this endpoint to finish registering a domain that is awaiting setup on your account.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --additional-details\n    {}\n\n  --domain-contacts\n    {\n      \"admin_id\": 546123,\n      \"billing_id\": 741288,\n      \"owner_id\": 741288,\n      \"tech_id\": 741288\n    }",
	Example: "  hostinger domains portfolio complete-setup mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := completeSetupBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsCompleteDomainSetupV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CompleteSetupCmd, utils.StringArg("domain"))
	CompleteSetupCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)")
	CompleteSetupCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON, @file or -; shape in the help above)")
	CompleteSetupCmd.RegisterFlagCompletionFunc("additional-details", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CompleteSetupCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func completeSetupBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		additionalDetailsVal, _ := cmd.Flags().GetString("additional-details")
		additionalDetailsValJSON, err := utils.JSONValue(additionalDetailsVal, "additional-details")
		if err != nil {
			return nil, err
		}
		body["additional_details"] = additionalDetailsValJSON
	}
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
