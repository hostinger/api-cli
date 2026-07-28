package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListDeliveryLogsCmd = &cobra.Command{
	Use:   "list-delivery-logs <order-id>",
	Short: "List webhook delivery logs",
	Long:  "Retrieve a paginated list of webhook delivery logs for the given mail\norder, including delivery outcome, duration, and retry counts.\nSupports filtering by mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailListWebhookDeliveryLogsV1WithResponse(context.TODO(), args[0], listDeliveryLogsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListDeliveryLogsCmd.Flags().StringP("mailbox-id", "", "", "Filter by the mailbox resource ID the webhooks are attached to")
	ListDeliveryLogsCmd.Flags().IntP("page", "", 0, "Page number")
	ListDeliveryLogsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listDeliveryLogsParams(cmd *cobra.Command) *client.MailListWebhookDeliveryLogsV1Params {
	params := &client.MailListWebhookDeliveryLogsV1Params{}
	if cmd.Flags().Changed("mailbox-id") {
		v, _ := cmd.Flags().GetString("mailbox-id")
		params.MailboxId = &v
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
