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

var ListWebsiteCmd = &cobra.Command{
	Use:     "list-website [domain]",
	Short:   "List website redirects",
	Long:    "Returns a paginated list of redirects configured for the selected website.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting redirects list-website mydomain.tld",
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
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		params := listWebsiteParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListWebsiteRedirectsV1WithResponse(ctx, username, domain, params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterScopedArgs(ListWebsiteCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	ListWebsiteCmd.Flags().IntP("page", "", 0, "Page number")
	ListWebsiteCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListWebsiteCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListWebsiteCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListWebsiteCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListWebsiteCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listWebsiteParams(cmd *cobra.Command) *client.HostingListWebsiteRedirectsV1Params {
	params := &client.HostingListWebsiteRedirectsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
