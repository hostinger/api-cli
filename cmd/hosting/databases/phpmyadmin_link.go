package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PhpmyadminLinkCmd = &cobra.Command{
	Use:     "phpmyadmin-link <username> <name>",
	Short:   "Get phpMyAdmin link",
	Long:    "Returns a direct sign-on link to phpMyAdmin for the specified database.\n\nUse this when a visual database interface is needed for SQL queries, imports, exports, or table management.\nThe database name must be the full name returned by the list databases endpoint.",
	Example: "  hostinger hosting databases phpmyadmin-link u123456789 u123456789_test_db",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetPhpMyAdminLinkV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(PhpmyadminLinkCmd, utils.StringArg("username"), utils.StringArg("name"))
}
