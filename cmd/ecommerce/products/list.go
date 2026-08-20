package products

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <store_id>",
	Short: "List products",
	Long:  "List a store's products newest first as lean summaries (name, status, thumbnail, variant\ncount and price range). Prices are integers in the smallest currency unit and live on\nvariants. Filter by status, free text or a set of product ids. Use include=variants to\nembed each product's variants with prices and inventory, and include=media to embed its media.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceListProductsV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringSliceP("product-ids", "", nil, "Restrict to these product ids. Doubles as a single-product lookup. Up to 200 ids.")
	ListCmd.Flags().StringSliceP("status", "", nil, "Product statuses to include. (one of: draft, proposed, published, rejected, archived)")
	ListCmd.Flags().StringP("q", "", "", "Free-text search over product title and SKU.")
	ListCmd.Flags().StringSliceP("include", "", nil, "Opt-in heavy data: \"variants\" embeds each product's variants; \"media\" embeds its media. (one of: variants, media)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
}

func listParams(cmd *cobra.Command) *client.EcommerceListProductsV1Params {
	params := &client.EcommerceListProductsV1Params{}
	if cmd.Flags().Changed("product-ids") {
		v, _ := cmd.Flags().GetStringSlice("product-ids")
		params.ProductIds = &v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetStringSlice("status")
		es := make([]client.EcommerceListProductsV1ParamsStatus, len(v))
		for i, s := range v {
			es[i] = client.EcommerceListProductsV1ParamsStatus(s)
		}
		params.Status = &es
	}
	if cmd.Flags().Changed("q") {
		v, _ := cmd.Flags().GetString("q")
		params.Q = &v
	}
	if cmd.Flags().Changed("include") {
		v, _ := cmd.Flags().GetStringSlice("include")
		es := make([]client.EcommerceListProductsV1ParamsInclude, len(v))
		for i, s := range v {
			es[i] = client.EcommerceListProductsV1ParamsInclude(s)
		}
		params.Include = &es
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
