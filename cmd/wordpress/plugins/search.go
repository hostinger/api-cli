package plugins

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:               "search",
	Short:             "Search WordPress plugins",
	Long:              "Search the WordPress.org plugin directory for plugins available to install.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.",
	Example:           "  hostinger wordpress plugins search --search seo",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := searchParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingSearchWordPressPluginsV1WithResponse(ctx, params)
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
	utils.RegisterArgs(SearchCmd)
	SearchCmd.Flags().StringP("search", "", "", "Search term to match against plugin names. Minimum 3 characters.")
	SearchCmd.MarkFlagRequired("search")
	SearchCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
}

func searchParams(cmd *cobra.Command) *client.HostingSearchWordPressPluginsV1Params {
	params := &client.HostingSearchWordPressPluginsV1Params{}
	searchVal, _ := cmd.Flags().GetString("search")
	params.Search = searchVal
	return params
}
