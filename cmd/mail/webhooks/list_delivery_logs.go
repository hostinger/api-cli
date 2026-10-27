package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListDeliveryLogsCmd = &cobra.Command{
	Use:     "list-delivery-logs [order-id]",
	Short:   "List webhook delivery logs",
	Long:    "Retrieve a paginated list of webhook delivery logs for the given mail\norder, including delivery outcome, duration, and retry counts.\nSupports filtering by mailbox.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail webhooks list-delivery-logs OR1a2b3c4d5e6f7g",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "id",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "mail")
		if err != nil {
			return api.CommandError(err)
		}
		params := listDeliveryLogsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListWebhookDeliveryLogsV1WithResponse(ctx, orderID, params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterScopedArgs(ListDeliveryLogsCmd, 1, utils.StringArg("order-id"))
	ListDeliveryLogsCmd.Flags().StringP("mailbox-id", "", "", "Filter by the mailbox resource ID the webhooks are attached to")
	ListDeliveryLogsCmd.Flags().IntP("page", "", 0, "Page number")
	ListDeliveryLogsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListDeliveryLogsCmd.RegisterFlagCompletionFunc("mailbox-id", cobra.NoFileCompletions)
	ListDeliveryLogsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListDeliveryLogsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListDeliveryLogsCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListDeliveryLogsCmd.MarkFlagsMutuallyExclusive("all", "page")
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
