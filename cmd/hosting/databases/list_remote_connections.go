package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListRemoteConnectionsCmd = &cobra.Command{
	Use:     "list-remote-connections [username]",
	Short:   "List database remote connections",
	Long:    "Returns the remote-access rules for the specified account: the remote hosts\n(IPv4/IPv6 addresses, or \"%\" for any host) allowed to connect to the account databases.\n\nUse the domain filter to only return rules for databases assigned to a specific domain.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting databases list-remote-connections u123456789",
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
		params := listRemoteConnectionsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListDatabaseRemoteConnectionsV1WithResponse(ctx, username, params)
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
	utils.RegisterScopedArgs(ListRemoteConnectionsCmd, 1, utils.StringArg("username"))
	ListRemoteConnectionsCmd.Flags().StringP("domain", "", "", "Filter remote connections by the domain the database is assigned to. Rules for databases not assigned to any domain are always included.")
	ListRemoteConnectionsCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func listRemoteConnectionsParams(cmd *cobra.Command) *client.HostingListDatabaseRemoteConnectionsV1Params {
	params := &client.HostingListDatabaseRemoteConnectionsV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
