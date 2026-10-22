package payments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListStoreProvidersCmd = &cobra.Command{
	Use:               "list-store-providers <store_id>",
	Short:             "List store payment providers",
	Long:              "List a store's payment providers, split into providers already connected to the store and\ngateways available to install. Never exposes gateway credentials, secrets, or configuration.",
	Example:           "  hostinger ecommerce payments list-store-providers store_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listStoreProvidersParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceListStorePaymentProvidersV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListStoreProvidersCmd, utils.StringArg("store_id"))
	ListStoreProvidersCmd.Flags().BoolP("include-currency-unsupported", "", false, "Include gateways that do not support the store currency in the available list.")
	ListStoreProvidersCmd.RegisterFlagCompletionFunc("include-currency-unsupported", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func listStoreProvidersParams(cmd *cobra.Command) *client.EcommerceListStorePaymentProvidersV1Params {
	params := &client.EcommerceListStorePaymentProvidersV1Params{}
	if cmd.Flags().Changed("include-currency-unsupported") {
		v, _ := cmd.Flags().GetBool("include-currency-unsupported")
		params.IncludeCurrencyUnsupported = &v
	}
	return params
}
