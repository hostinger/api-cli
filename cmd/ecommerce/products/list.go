package products

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <store_id>",
	Short:             "List products",
	Long:              "List a store's products newest first as lean summaries (name, status, thumbnail, variant\ncount and price range). Prices are integers in the smallest currency unit and live on\nvariants. Filter by status, free text or a set of product ids. Use include=variants to\nembed each product's variants with prices and inventory, and include=media to embed its media.",
	Example:           "  hostinger ecommerce products list store_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceListProductsV1WithResponse(ctx, args[0], params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(ListCmd, utils.StringArg("store_id"))
	ListCmd.Flags().StringSliceP("product-ids", "", nil, "Restrict to these product ids. Doubles as a single-product lookup. Up to 200 ids.")
	ListCmd.Flags().StringSliceP("status", "", nil, "Product statuses to include. (one of: draft, proposed, published, rejected, archived)")
	ListCmd.Flags().StringP("q", "", "", "Free-text search over product title and SKU.")
	ListCmd.Flags().StringSliceP("include", "", nil, "Opt-in heavy data: \"variants\" embeds each product's variants; \"media\" embeds its media. (one of: variants, media)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.RegisterFlagCompletionFunc("product-ids", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"draft", "proposed", "published", "rejected", "archived"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("q", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("include", cobra.FixedCompletions([]cobra.Completion{"variants", "media"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
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
