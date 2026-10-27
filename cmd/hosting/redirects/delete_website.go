package redirects

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteCmd = &cobra.Command{
	Use:     "delete-website [domain]",
	Short:   "Delete website redirect",
	Long:    "Permanently deletes the redirect identified by its source URL.\n\nPass the `from` value exactly as returned by the list redirects endpoint.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting redirects delete-website mydomain.tld --from https://example.com/old-page",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 0
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		params := deleteWebsiteParams(cmd)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Delete website redirect", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWebsiteRedirectV1WithResponse(ctx, username, domain, params)
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
	utils.RegisterScopedArgs(DeleteWebsiteCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	DeleteWebsiteCmd.Flags().StringP("from", "", "", "Source URL returned by the list redirects endpoint.")
	DeleteWebsiteCmd.MarkFlagRequired("from")
	DeleteWebsiteCmd.RegisterFlagCompletionFunc("from", cobra.NoFileCompletions)
	DeleteWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteWebsiteParams(cmd *cobra.Command) *client.HostingDeleteWebsiteRedirectV1Params {
	params := &client.HostingDeleteWebsiteRedirectV1Params{}
	fromVal, _ := cmd.Flags().GetString("from")
	params.From = fromVal
	return params
}
