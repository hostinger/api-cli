package subscriptions

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableAutoRenewalCmd = &cobra.Command{
	Use:   "enable-auto-renewal <subscription-id>",
	Short: "Enable auto-renewal",
	Long:  "Enable auto-renewal for a subscription.\n\nUse this endpoint when enable auto-renewal for a subscription.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingEnableAutoRenewalV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
