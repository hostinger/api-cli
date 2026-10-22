package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Get catalog item list",
	Long:              "Retrieve catalog items available for order.\n\nPrices in catalog items is displayed as cents (without floating point),\ne.g: float `17.99` is displayed as integer `1799`.\n\nUse this endpoint to view available services and pricing before placing orders.",
	Example:           "  hostinger billing catalog list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "category", []string{"DOMAIN", "VPS", "EMAIL", "HOSTING"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().BillingGetCatalogItemListV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().StringP("category", "", "", "Filter catalog items by category (one of: DOMAIN, VPS, EMAIL, HOSTING)")
	ListCmd.Flags().StringP("name", "", "", "Filter catalog items by name. Use `*` for wildcard search, e.g. `.COM*` to find .com domain")
	ListCmd.RegisterFlagCompletionFunc("category", cobra.FixedCompletions([]cobra.Completion{"DOMAIN", "VPS", "EMAIL", "HOSTING"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func listParams(cmd *cobra.Command) *client.BillingGetCatalogItemListV1Params {
	params := &client.BillingGetCatalogItemListV1Params{}
	if cmd.Flags().Changed("category") {
		v, _ := cmd.Flags().GetString("category")
		e := client.BillingGetCatalogItemListV1ParamsCategory(v)
		params.Category = &e
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		params.Name = &v
	}
	return params
}
