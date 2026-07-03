package stores

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var MetadataCmd = &cobra.Command{
	Use:   "metadata <store_id>",
	Short: "Get store metadata",
	Long:  "Get a store's readiness metadata: whether payment methods and shipping are configured,\nplus its default currency. Useful to verify prerequisites before building a storefront.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceGetStoreMetadataV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
