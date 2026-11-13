package pocketbase

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListBackendsCmd = &cobra.Command{
	Use:     "list-backends [username]",
	Short:   "List account PocketBase backends",
	Long:    "Lists the PocketBase backends of all websites of a hosting account with their URLs, live\nstate, version and disk usage. A PocketBase backend gives a website a database,\nauthentication, file storage and a realtime REST API, served from the website's own\ndomain. An empty list means no website of the account has a PocketBase backend.\n\nStates: `asleep` is normal (the instance sleeps when idle and wakes on the first request);\n`starting`, `running` and `stopping` are healthy.\n\n`failed` means the last start did not succeed. Hostinger starts it again on its own after\n`retry_after_seconds`, and the data is intact: wait and call again, do not delete the\nbackend. `failure_reason` says why: `process_limit` means the hosting account has no free\nprocess slots, usually because its Node.js apps are running.\n\n`unregistered` means the instance is broken. `missing` means its files are gone from the\nserver, and creating the PocketBase backend again sets up a new, empty one. `state` is null\nwhen the server could not be reached; retry later.\n\nPocketBase is rolling out gradually: an account without access yet gets 422.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting pocketbase list-backends u123456789",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "username",
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
		username, args, err := site.ResolveUsername(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListAccountPocketBaseBackendsV1WithResponse(ctx, username)
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
	utils.RegisterScopedArgs(ListBackendsCmd, 1, utils.StringArg("username"))
}
