package product_variants

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <store_id> <product_id> <variant_id>",
	Short: "Delete a product variant",
	Long:  "Delete a single variant from the product.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceDeleteAProductVariantV1WithResponse(context.TODO(), args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
