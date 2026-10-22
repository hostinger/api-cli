package miscellaneous

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CustomStorefrontSetupInstructionsCmd = &cobra.Command{
	Use:               "custom-storefront-setup-instructions",
	Short:             "Get custom storefront setup instructions",
	Long:              "Retrieve step-by-step setup instructions, formatted as Markdown, for connecting a custom sales\nchannel to your store and keeping your catalog, orders, shipping and payments in sync through\nthe Ecommerce API.",
	Example:           "  hostinger ecommerce miscellaneous custom-storefront-setup-instructions",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceGetCustomStorefrontSetupInstructionsV1WithResponse(ctx)
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
	utils.RegisterArgs(CustomStorefrontSetupInstructionsCmd)
}
