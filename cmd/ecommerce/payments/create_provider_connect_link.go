package payments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateProviderConnectLinkCmd = &cobra.Command{
	Use:               "create-provider-connect-link <store_id> <provider_id>",
	Short:             "Create a payment provider connect link",
	Long:              "Create an onboarding link for connecting a payment gateway to the store. Returns the gateway\nonboarding URL for the merchant to open and a deep-link into the store admin.",
	Example:           "  hostinger ecommerce payments create-provider-connect-link store_01J8Z5F8W9K8M4A7B3C2D1E0FG stripe",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateAPaymentProviderConnectLinkV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(CreateProviderConnectLinkCmd, utils.StringArg("store_id"), utils.StringArg("provider_id"))
}
