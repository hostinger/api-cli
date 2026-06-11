package catalog

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get catalog item list",
	Long: `This endpoint retrieves a list of available catalog items that can be ordered.

Prices in the response are displayed in cents.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingGetCatalogItemListV1WithResponse(context.TODO(), catalogListRequestParameters(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("category", "", "", "Filter catalog items by category (one of: DOMAIN, VPS)")
	ListCmd.Flags().StringP("name", "", "", "Filter catalog items by name. Use * for wildcard search, e.g. .COM* to find .com domain")
}

func catalogListRequestParameters(cmd *cobra.Command) *client.BillingGetCatalogItemListV1Params {
	category, _ := cmd.Flags().GetString("category")
	name, _ := cmd.Flags().GetString("name")

	params := &client.BillingGetCatalogItemListV1Params{
		Name: utils.StringPtrOrNil(name),
	}

	if category != "" {
		c := client.BillingGetCatalogItemListV1ParamsCategory(category)
		params.Category = &c
	}

	return params
}
