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

var StartCmd = &cobra.Command{
	Use:               "start",
	Short:             "Start domain transfer",
	Long:              "Transfer a domain from another registrar to your account.\n\nThe transfer runs on a domain transfer service you have already purchased.\n\nBefore making request, unlock the domain at the current registrar and get its authorization\ncode.\n\nA successful response means the transfer has been started. Completion depends on the current\nregistrar and can be followed with the transfer list endpoint (#tag/domains-transfer).\n\nIf no WHOIS information is provided, default contact information for that TLD will be used.\nBefore making request, ensure WHOIS information for desired TLD exists in your account.\n\nUse this endpoint to bring domains registered elsewhere into your account.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --domain-contacts\n    {\n      \"admin_id\": 546123,\n      \"billing_id\": 741288,\n      \"owner_id\": 741288,\n      \"tech_id\": 741288\n    }",
	Example:           "  hostinger domains transfer start --auth-code Auth123Code456 --domain my-new-domain.tld",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := startBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsStartDomainTransferV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(StartCmd)
	StartCmd.Flags().StringP("auth-code", "", "", "Authorization code from the current registrar")
	StartCmd.Flags().StringP("domain", "", "", "Domain name")
	StartCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON, @file or -; shape in the help above)")
	StartCmd.Flags().BoolP("should-keep-ns", "", true, "Keep the existing nameservers of the domain")
	StartCmd.MarkFlagRequired("auth-code")
	StartCmd.MarkFlagRequired("domain")
	StartCmd.RegisterFlagCompletionFunc("auth-code", cobra.NoFileCompletions)
	StartCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	StartCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	StartCmd.RegisterFlagCompletionFunc("should-keep-ns", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func startBody(cmd *cobra.Command) (map[string]any, error) {
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
