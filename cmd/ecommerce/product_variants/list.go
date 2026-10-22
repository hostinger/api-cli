package product_variants

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <store_id> <product_id>",
	Short:             "List product variants",
	Long:              "List a product's variants, ordered by rank, with their options, prices and inventory.\nPrices are integers in the smallest currency unit and live on variants.",
	Example:           "  hostinger ecommerce product-variants list store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceListProductVariantsV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.EcommerceListProductVariantsV1Params {
	params := &client.EcommerceListProductVariantsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
