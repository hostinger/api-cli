package catalog

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get catalog item list",
	Long:  "Retrieve catalog items available for order.\n\nPrices in catalog items is displayed as cents (without floating point),\ne.g: float `17.99` is displayed as integer `1799`.\n\nUse this endpoint to view available services and pricing before placing orders.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "category", []string{"DOMAIN", "VPS"})
		r, err := api.Request().BillingGetCatalogItemListV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("category", "", "", "Filter catalog items by category (one of: DOMAIN, VPS)")
	ListCmd.Flags().StringP("name", "", "", "Filter catalog items by name. Use `*` for wildcard search, e.g. `.COM*` to find .com domain")
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
