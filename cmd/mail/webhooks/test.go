package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test <webhook-id>",
	Short: "Test webhook",
	Long:  "Send a test delivery to the webhook URL and return the result. Test\nrequests are rate limited upstream.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailTestWebhookV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
