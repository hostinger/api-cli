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

var DeleteRemoteConnectionCmd = &cobra.Command{
	Use:     "delete-remote-connection [username] <name>",
	Short:   "Delete database remote connection",
	Long:    "Permanently removes a remote-access rule, revoking the given host's remote access to the database.\n\nIdentify the rule with the required ip query parameter (the IPv4/IPv6 address, or \"%\",\nexactly as returned by the list remote connections endpoint). The database name must be\nthe full name returned by the list databases endpoint.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting databases delete-remote-connection u123456789 u123456789_test_db --ip 192.0.2.10",
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
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 1
		username, args, err := site.ResolveUsername(cmd, args, 1)
		if err != nil {
			return api.CommandError(err)
		}
		params := deleteRemoteConnectionParams(cmd)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Delete database remote connection", append([]string{username}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteDatabaseRemoteConnectionV1WithResponse(ctx, username, args[0], params)
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
	utils.RegisterScopedArgs(DeleteRemoteConnectionCmd, 1, utils.StringArg("username"), utils.StringArg("name"))
	DeleteRemoteConnectionCmd.Flags().StringP("ip", "", "", "Remote host to revoke: the IPv4/IPv6 address, or \"%\", exactly as returned by the list remote connections endpoint.")
	DeleteRemoteConnectionCmd.MarkFlagRequired("ip")
	DeleteRemoteConnectionCmd.RegisterFlagCompletionFunc("ip", cobra.NoFileCompletions)
	DeleteRemoteConnectionCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteRemoteConnectionParams(cmd *cobra.Command) *client.HostingDeleteDatabaseRemoteConnectionV1Params {
	params := &client.HostingDeleteDatabaseRemoteConnectionV1Params{}
	ipVal, _ := cmd.Flags().GetString("ip")
	params.Ip = ipVal
	return params
}
