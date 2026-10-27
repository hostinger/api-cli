package integrations

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListConnectedWebsitesCmd = &cobra.Command{
	Use:               "list-connected-websites <profile-uuid>",
	Short:             "List connected websites",
	Long:              "Get the websites connected to the profile, each with its `domain` and `platform_type`.\n\nA domain is connected to at most one profile, so to find which plan a site is on, call this\nfor each profile. Only connected sites are listed, not every website the customer owns, and\neCommerce stores are not included - use the list eCommerce stores endpoint for those. Pass\n`domain` to the disconnect WordPress site or disconnect AI Website Builder site endpoint.",
	Example:           "  hostinger reach integrations list-connected-websites 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListConnectedWebsitesV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListConnectedWebsitesCmd, utils.StringArg("profile-uuid"))
}
