package products

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateImageUploadUrlCmd = &cobra.Command{
	Use:               "create-image-upload-url <store_id> <product_id>",
	Short:             "Create a product image upload URL",
	Long:              "Returns a signed URL to upload a product image to (multipart/form-data POST). Then call the\nattach-image endpoint with the returned object_name to scan and attach it to the product.",
	Example:           "  hostinger ecommerce products create-image-upload-url store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateAProductImageUploadURLV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(CreateImageUploadUrlCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
}
