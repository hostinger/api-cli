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

var StartOutgoingCmd = &cobra.Command{
	Use:     "start-outgoing [domain]",
	Short:   "Start outgoing domain move",
	Long:    "Initiate a move of a specified domain to another Hostinger account.\n\nThe receiving account has to already exist and accept the move before the domain changes hands.\n\nThe domain must be active. The subscription it belongs to is resolved automatically,\nand the request is rejected with a 404 status code when the domain has no domain\nsubscription of its own.\n\nDomains protected by premium protection require an additional verification step,\nsuch requests are rejected with a 428 status code.\n\nUse this endpoint to hand a domain over to another Hostinger user.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains move start-outgoing mydomain.tld --new-customer-email new-owner@example.com",
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
		body, err := startOutgoingBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsStartOutgoingDomainMoveV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(StartOutgoingCmd, 1, utils.StringArg("domain"))
	StartOutgoingCmd.Flags().StringP("new-customer-email", "", "", "Email address of the Hostinger account receiving the domain")
	StartOutgoingCmd.MarkFlagRequired("new-customer-email")
	StartOutgoingCmd.RegisterFlagCompletionFunc("new-customer-email", cobra.NoFileCompletions)
}

func startOutgoingBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	newCustomerEmailVal, _ := cmd.Flags().GetString("new-customer-email")
	body["new_customer_email"] = newCustomerEmailVal
	return body, nil
}
