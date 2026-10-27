package portfolio

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

var UpdateNameserversCmd = &cobra.Command{
	Use:     "update-nameservers [domain]",
	Short:   "Update domain nameservers",
	Long:    "Set nameservers for a specified domain.\n\nBe aware, that improper nameserver configuration can lead to the domain being unresolvable or unavailable.\n\nUse this endpoint to configure custom DNS hosting for domains.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains portfolio update-nameservers mydomain.tld --ns1 ns1.some-nameserver.tld --ns2 ns2.some-nameserver.tld",
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
		body, err := updateNameserversBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsUpdateDomainNameserversV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(UpdateNameserversCmd, 1, utils.StringArg("domain"))
	UpdateNameserversCmd.Flags().StringP("ns1", "", "", "First name server")
	UpdateNameserversCmd.Flags().StringP("ns2", "", "", "Second name server")
	UpdateNameserversCmd.Flags().StringP("ns3", "", "", "Third name server")
	UpdateNameserversCmd.Flags().StringP("ns4", "", "", "Fourth name server")
	UpdateNameserversCmd.MarkFlagRequired("ns1")
	UpdateNameserversCmd.MarkFlagRequired("ns2")
	UpdateNameserversCmd.RegisterFlagCompletionFunc("ns1", cobra.NoFileCompletions)
	UpdateNameserversCmd.RegisterFlagCompletionFunc("ns2", cobra.NoFileCompletions)
	UpdateNameserversCmd.RegisterFlagCompletionFunc("ns3", cobra.NoFileCompletions)
	UpdateNameserversCmd.RegisterFlagCompletionFunc("ns4", cobra.NoFileCompletions)
}

func updateNameserversBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	ns1Val, _ := cmd.Flags().GetString("ns1")
	body["ns1"] = ns1Val
	ns2Val, _ := cmd.Flags().GetString("ns2")
	body["ns2"] = ns2Val
	if cmd.Flags().Changed("ns3") {
		v, _ := cmd.Flags().GetString("ns3")
		body["ns3"] = v
	}
	if cmd.Flags().Changed("ns4") {
		v, _ := cmd.Flags().GetString("ns4")
		body["ns4"] = v
	}
	return body, nil
}
