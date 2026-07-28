package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <order-id>",
	Short: "List webhooks",
	Long:  "Retrieve a paginated list of webhooks belonging to the given mail\norder. Supports filtering by mailbox and status. The webhook secret\nis never included; it is returned only when a webhook is created or\nits secret is regenerated.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"})
		r, err := api.Request().MailListWebhooksV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("mailbox-id", "", "", "Filter by the mailbox resource ID the webhooks are attached to")
	ListCmd.Flags().StringP("status", "", "", "Filter webhooks by status (one of: active, disabled, paused)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.MailListWebhooksV1Params {
	params := &client.MailListWebhooksV1Params{}
	if cmd.Flags().Changed("mailbox-id") {
		v, _ := cmd.Flags().GetString("mailbox-id")
		params.MailboxId = &v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.MailListWebhooksV1ParamsStatus(v)
		params.Status = &e
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
