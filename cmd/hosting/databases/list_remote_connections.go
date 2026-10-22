package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListRemoteConnectionsCmd = &cobra.Command{
	Use:     "list-remote-connections <username>",
	Short:   "List database remote connections",
	Long:    "Returns the remote-access rules for the specified account: the remote hosts\n(IPv4/IPv6 addresses, or \"%\" for any host) allowed to connect to the account databases.\n\nUse the domain filter to only return rules for databases assigned to a specific domain.",
	Example: "  hostinger hosting databases list-remote-connections u123456789",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listRemoteConnectionsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListDatabaseRemoteConnectionsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListRemoteConnectionsCmd, utils.StringArg("username"))
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
