package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListBuildsCmd = &cobra.Command{
	Use:     "list-builds [domain]",
	Short:   "List NodeJS builds",
	Long:    "Retrieve a paginated list of Node.js build processes for a specific website.\n\nEach build represents a single run of the Node.js build pipeline. Use the `states`\nquery parameter to filter results by build state (pending, running, completed, failed).\nUse the `uuid` from a build to poll its output via the `Get Node.js Build Logs` endpoint.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting nodejs list-builds mydomain.tld",
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
		params := listBuildsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListNodeJSBuildsV1WithResponse(ctx, username, domain, params)
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
	utils.RegisterScopedArgs(ListBuildsCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	ListBuildsCmd.Flags().IntP("page", "", 0, "Page number")
	ListBuildsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListBuildsCmd.Flags().StringSliceP("states", "", nil, "Build states to filter by (one of: pending, running, completed, failed)")
	ListBuildsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListBuildsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListBuildsCmd.RegisterFlagCompletionFunc("states", cobra.FixedCompletions([]cobra.Completion{"pending", "running", "completed", "failed"}, cobra.ShellCompDirectiveNoFileComp))
	ListBuildsCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListBuildsCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listBuildsParams(cmd *cobra.Command) *client.HostingListNodeJSBuildsV1Params {
	params := &client.HostingListNodeJSBuildsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("states") {
		v, _ := cmd.Flags().GetStringSlice("states")
		es := make([]client.HostingListNodeJSBuildsV1ParamsStates, len(v))
		for i, s := range v {
			es[i] = client.HostingListNodeJSBuildsV1ParamsStates(s)
		}
		params.States = &es
	}
	return params
}
