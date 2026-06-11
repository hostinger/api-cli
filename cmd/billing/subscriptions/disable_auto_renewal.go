package subscriptions

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var DisableAutoRenewalCmd = &cobra.Command{
	Use:   "disable-auto-renewal <subscription ID>",
	Short: "Disable subscription auto-renewal",
	Long:  `This endpoint disables auto-renewal for a subscription, which stops any further auto-renewal of the associated service.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingDisableAutoRenewalV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
