package integrations

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListEcommerceStoresCmd = &cobra.Command{
	Use:               "list-ecommerce-stores <profile-uuid>",
	Short:             "List Ecommerce stores",
	Long:              "Get all of the customer's Hostinger eCommerce stores, each with its `store_id` and\n`connection_status` for the profile in the path.\n\nA store is connected to at most one profile. `connected_to_other_profile` does not name that\nprofile - call this for each profile to find it. The list is empty, not an error, when the\neCommerce service cannot be reached. Pass `store_id` to the connect or disconnect eCommerce\nstore endpoint.",
	Example:           "  hostinger reach integrations list-ecommerce-stores 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListEcommerceStoresV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListEcommerceStoresCmd, utils.StringArg("profile-uuid"))
}
