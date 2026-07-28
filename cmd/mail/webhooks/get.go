package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <webhook-id>",
	Short: "Get webhook",
	Long:  "Retrieve the details of a single webhook. The webhook secret is never\nincluded; it is returned only when a webhook is created or its secret\nis regenerated.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailGetWebhookV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
