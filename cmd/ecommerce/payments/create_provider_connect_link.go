package payments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateProviderConnectLinkCmd = &cobra.Command{
	Use:   "create-provider-connect-link <store_id> <provider_id>",
	Short: "Create a payment provider connect link",
	Long:  "Create an onboarding link for connecting a payment gateway to the store. Returns the gateway\nonboarding URL for the merchant to open and a deep-link into the store admin.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().EcommerceCreateAPaymentProviderConnectLinkV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
