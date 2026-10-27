package integrations

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisconnectEcommerceStoreCmd = &cobra.Command{
	Use:               "disconnect-ecommerce-store <profile-uuid> <store-id>",
	Short:             "Disconnect Ecommerce store",
	Long:              "Disconnect an eCommerce store from the profile.\n\nReach stops receiving the store's customers and orders, and its products and coupons are\nremoved from Reach; contacts and past orders stay. Take `store_id` from the list eCommerce\nstores endpoint; for a store that is `connected_to_other_profile`, call this with that\nprofile. Also succeeds when the store is not connected to this profile, so success is not\nproof that anything was disconnected. Confirm the store and the plan with the customer\nbefore calling.",
	Example:           "  hostinger reach integrations disconnect-ecommerce-store 550e8400-e09b-41d4-a716-400055000000 ec_123",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Disconnect Ecommerce store", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDisconnectEcommerceStoreV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(DisconnectEcommerceStoreCmd, utils.StringArg("profile-uuid"), utils.StringArg("store-id"))
	DisconnectEcommerceStoreCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
