package subscriptions

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var EnableAutoRenewalCmd = &cobra.Command{
	Use:   "enable-auto-renewal <subscription ID>",
	Short: "Enable subscription auto-renewal",
	Long:  `This endpoint enables auto-renewal for a subscription, so the associated service is renewed automatically.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingEnableAutoRenewalV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
