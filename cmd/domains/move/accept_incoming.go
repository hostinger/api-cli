package move

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AcceptIncomingCmd = &cobra.Command{
	Use:     "accept-incoming [domain]",
	Short:   "Accept incoming domain move",
	Long:    "Accept an incoming move for a specified domain.\n\nThe provided WHOIS profiles become the contacts of the domain, so they must belong\nto your account and satisfy the requirements of the TLD. Only the contact types the\ndomain actually uses are applied, but all four profile IDs have to be provided.\n\nThe move has to still be waiting for your decision, already accepted moves\ncannot be accepted again.\n\nAccepting does not complete the move. A confirmation email is sent to the email address of\nthe new owner contact, and the domain changes hands only after the change is confirmed from it.\nUntil then the move stays in the `activating` status, which can be followed with the\nincoming move endpoint (#tag/domains-move).\n\nUse this endpoint to take ownership of a domain offered to you.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --domain-contacts\n    {\n      \"admin_id\": 114698,\n      \"billing_id\": 154698,\n      \"owner_id\": 614698,\n      \"tech_id\": 524248\n    }",
	Example: "  hostinger domains move accept-incoming mydomain.tld --domain-contacts @domain-contacts.json",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		domain, args, err := site.ResolveDomain(cmd, args, 0, "domains")
		if err != nil {
			return api.CommandError(err)
		}
		body, err := acceptIncomingBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsAcceptIncomingDomainMoveV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(AcceptIncomingCmd, 1, utils.StringArg("domain"))
	AcceptIncomingCmd.Flags().StringP("domain-contacts", "", "", "WHOIS profiles of the accepting account. Only the contact types required by the TLD are applied, but all four IDs must be provided. (JSON, @file or -; shape in the help above)")
	AcceptIncomingCmd.MarkFlagRequired("domain-contacts")
	AcceptIncomingCmd.RegisterFlagCompletionFunc("domain-contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func acceptIncomingBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainContactsVal, _ := cmd.Flags().GetString("domain-contacts")
	domainContactsValJSON, err := utils.JSONValue(domainContactsVal, "domain-contacts")
	if err != nil {
		return nil, err
	}
	body["domain_contacts"] = domainContactsValJSON
	return body, nil
}
