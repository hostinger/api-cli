package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <webhook-id>",
	Short: "Delete webhook",
	Long:  "Permanently delete a webhook. This action cannot be undone. After\ndeletion the URL no longer receives event notifications.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteWebhookV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
