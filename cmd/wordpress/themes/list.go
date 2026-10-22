package themes

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List WordPress themes",
	Long:              "List WordPress themes available to install.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/themes/install.",
	Example:           "  hostinger wordpress themes list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListWordPressThemesV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().IntP("order-id", "", 0, "Optionally scope themes to a specific order.")
	ListCmd.Flags().StringP("search", "", "", "Search term to match against theme names.")
	ListCmd.RegisterFlagCompletionFunc("order-id", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
}

func listParams(cmd *cobra.Command) *client.HostingListWordPressThemesV1Params {
	params := &client.HostingListWordPressThemesV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	return params
}
