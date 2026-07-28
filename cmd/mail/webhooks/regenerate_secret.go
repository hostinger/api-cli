package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RegenerateSecretCmd = &cobra.Command{
	Use:   "regenerate-secret <webhook-id>",
	Short: "Regenerate webhook secret",
	Long:  "Regenerate the secret of a webhook. The previous secret is\nimmediately invalidated. The new secret is returned only in this\nresponse and is sent as a bearer token with every delivery.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailRegenerateWebhookSecretV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
