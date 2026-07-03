package miscellaneous

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CustomStorefrontSetupInstructionsCmd = &cobra.Command{
	Use:   "custom-storefront-setup-instructions",
	Short: "Get custom storefront setup instructions",
	Long:  "Retrieve step-by-step setup instructions, formatted as Markdown, for connecting a custom sales\nchannel to your store and keeping your catalog, orders, shipping and payments in sync through\nthe Ecommerce API.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceGetCustomStorefrontSetupInstructionsV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
