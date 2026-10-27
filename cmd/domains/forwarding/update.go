package forwarding

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

var UpdateCmd = &cobra.Command{
	Use:     "update [domain]",
	Short:   "Update domain forwarding",
	Long:    "Update domain forwarding configuration.\n\nUse this endpoint to modify existing redirect configuration for domains.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains forwarding update mydomain.tld --redirect-type 301 --redirect-url https://forward.to.my.url",
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
		if err := utils.EnumCheck(cmd, "redirect-type", []string{"301", "302"}); err != nil {
			return err
		}
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsUpdateDomainForwardingV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(UpdateCmd, 1, utils.StringArg("domain"))
	UpdateCmd.Flags().StringP("redirect-type", "", "", "Redirect type (one of: 301, 302)")
	UpdateCmd.Flags().StringP("redirect-url", "", "", "URL to forward domain to")
	UpdateCmd.MarkFlagRequired("redirect-type")
	UpdateCmd.MarkFlagRequired("redirect-url")
	UpdateCmd.RegisterFlagCompletionFunc("redirect-type", cobra.FixedCompletions([]cobra.Completion{"301", "302"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("redirect-url", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	redirectTypeVal, _ := cmd.Flags().GetString("redirect-type")
	body["redirect_type"] = redirectTypeVal
	redirectUrlVal, _ := cmd.Flags().GetString("redirect-url")
	body["redirect_url"] = redirectUrlVal
	return body, nil
}
