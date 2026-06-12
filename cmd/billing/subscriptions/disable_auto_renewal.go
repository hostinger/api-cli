package subscriptions

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisableAutoRenewalCmd = &cobra.Command{
	Use:   "disable-auto-renewal <subscription-id>",
	Short: "Disable auto-renewal",
	Long:  "Disable auto-renewal for a subscription.\n\nUse this endpoint when disable auto-renewal for a subscription.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingDisableAutoRenewalV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
