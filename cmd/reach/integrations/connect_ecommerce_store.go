package integrations

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ConnectEcommerceStoreCmd = &cobra.Command{
	Use:               "connect-ecommerce-store <profile-uuid> <store-id>",
	Short:             "Connect Ecommerce store",
	Long:              "Connect one of the customer's eCommerce stores to the profile, so Reach receives its\ncustomers and orders.\n\nTake `store_id` from the list eCommerce stores endpoint. Fails with not found\n(`ECOMMERCE_STORE_NOT_FOUND`) when the store is not one of the customer's own stores, also\nwhen the eCommerce service cannot be reached. Fails with a conflict\n(`ECOMMERCE_CONNECTION_ALREADY_EXISTS`) when the store is connected to any profile or this\nprofile already has a store - to move a store to another plan, disconnect it from the\nprofile that holds it first.",
	Example:           "  hostinger reach integrations connect-ecommerce-store 550e8400-e09b-41d4-a716-400055000000 ec_123",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachConnectEcommerceStoreV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(ConnectEcommerceStoreCmd, utils.StringArg("profile-uuid"), utils.StringArg("store-id"))
}
