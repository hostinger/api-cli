package plugins

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListSuggestedCmd = &cobra.Command{
	Use:               "list-suggested",
	Short:             "List suggested WordPress plugins",
	Long:              "List curated plugin suggestions grouped by website type.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.",
	Example:           "  hostinger wordpress plugins list-suggested",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listSuggestedParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListSuggestedWordPressPluginsV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListSuggestedCmd)
	ListSuggestedCmd.Flags().IntP("order-id", "", 0, "Optionally scope suggestions to a specific order.")
	ListSuggestedCmd.RegisterFlagCompletionFunc("order-id", cobra.NoFileCompletions)
}

func listSuggestedParams(cmd *cobra.Command) *client.HostingListSuggestedWordPressPluginsV1Params {
	params := &client.HostingListSuggestedWordPressPluginsV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	return params
}
