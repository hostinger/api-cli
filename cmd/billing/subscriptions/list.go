package subscriptions

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get subscription list",
	Long:  "Retrieve a list of all subscriptions associated with your account.\n\nUse this endpoint to monitor active services and billing status.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().BillingGetSubscriptionListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
