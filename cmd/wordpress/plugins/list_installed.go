package plugins

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListInstalledCmd = &cobra.Command{
	Use:     "list-installed [username] <software>",
	Short:   "List installed WordPress plugins",
	Long:    "List plugins installed on a WordPress installation, including their status,\navailable updates and known vulnerabilities.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress plugins list-installed u123456789 1232456789",
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
		username, args, err := site.ResolveUsername(cmd, args, 1)
		if err != nil {
			return api.CommandError(err)
		}
		if err := utils.EnumCheck(cmd, "category", []string{"cache"}); err != nil {
			return err
		}
		params := listInstalledParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListInstalledWordPressPluginsV1WithResponse(ctx, username, args[0], params)
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
	utils.RegisterScopedArgs(ListInstalledCmd, 1, utils.StringArg("username"), utils.StringArg("software"))
	ListInstalledCmd.Flags().StringP("category", "", "", "Filter installed plugins by category. (one of: cache)")
	ListInstalledCmd.RegisterFlagCompletionFunc("category", cobra.FixedCompletions([]cobra.Completion{"cache"}, cobra.ShellCompDirectiveNoFileComp))
}

func listInstalledParams(cmd *cobra.Command) *client.HostingListInstalledWordPressPluginsV1Params {
	params := &client.HostingListInstalledWordPressPluginsV1Params{}
	if cmd.Flags().Changed("category") {
		v, _ := cmd.Flags().GetString("category")
		e := client.HostingListInstalledWordPressPluginsV1ParamsCategory(v)
		params.Category = &e
	}
	return params
}
