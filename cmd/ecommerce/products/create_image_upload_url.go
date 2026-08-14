package products

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateImageUploadUrlCmd = &cobra.Command{
	Use:   "create-image-upload-url <store_id> <product_id>",
	Short: "Create a product image upload URL",
	Long:  "Returns a signed URL to upload a product image to (multipart/form-data POST). Then call the\nattach-image endpoint with the returned object_name to scan and attach it to the product.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceCreateAProductImageUploadURLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
