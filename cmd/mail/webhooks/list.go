package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list <order-id>",
	Short:   "List webhooks",
	Long:    "Retrieve a paginated list of webhooks belonging to the given mail\norder. Supports filtering by mailbox and status. The webhook secret\nis never included; it is returned only when a webhook is created or\nits secret is regenerated.",
	Example: "  hostinger mail webhooks list OR1a2b3c4d5e6f7g",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListWebhooksV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("order-id"))
	ListCmd.Flags().StringP("mailbox-id", "", "", "Filter by the mailbox resource ID the webhooks are attached to")
	ListCmd.Flags().StringP("status", "", "", "Filter webhooks by status (one of: active, disabled, paused)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.RegisterFlagCompletionFunc("mailbox-id", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "disabled", "paused"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
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
