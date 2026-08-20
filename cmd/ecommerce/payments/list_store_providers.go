package payments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListStoreProvidersCmd = &cobra.Command{
	Use:   "list-store-providers <store_id>",
	Short: "List store payment providers",
	Long:  "List a store's payment providers, split into providers already connected to the store and\ngateways available to install. Never exposes gateway credentials, secrets, or configuration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceListStorePaymentProvidersV1WithResponse(context.TODO(), args[0], listStoreProvidersParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListStoreProvidersCmd.Flags().BoolP("include-currency-unsupported", "", false, "Include gateways that do not support the store currency in the available list.")
}

func listStoreProvidersParams(cmd *cobra.Command) *client.EcommerceListStorePaymentProvidersV1Params {
	params := &client.EcommerceListStorePaymentProvidersV1Params{}
	if cmd.Flags().Changed("include-currency-unsupported") {
		v, _ := cmd.Flags().GetBool("include-currency-unsupported")
		params.IncludeCurrencyUnsupported = &v
	}
	return params
}
