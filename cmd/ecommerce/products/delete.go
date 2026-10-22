package products

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:               "delete <store_id> <product_id>",
	Short:             "Delete a product",
	Long:              "Delete a product and its variants from the store. A subscription product with active\nsubscribers is archived instead of deleted so its data stays available.",
	Example:           "  hostinger ecommerce products delete store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete a product", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceDeleteAProductV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
