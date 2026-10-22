package logs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListMailboxActionCmd = &cobra.Command{
	Use:     "list-mailbox-action <order-id>",
	Short:   "List mailbox action logs",
	Long:    "Retrieve paginated mailbox action logs (message and mailbox events)\nfor a mailbox in the given mail order. The mailbox email must belong\nto the order's domain. Supports date range and event type filters.\nResults are sorted by timestamp descending.",
	Example: "  hostinger mail logs list-mailbox-action OR1a2b3c4d5e6f7g --email user@example.com",
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
		if err := utils.EnumCheck(cmd, "event", []string{"MessageNew", "MessageRead", "MessageAppend", "MessageExpunge", "MailboxCreate", "MailboxDelete", "MailboxRename"}); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "email", "email"); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "date", "date"); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "from-date", "date-time"); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "to-date", "date-time"); err != nil {
			return err
		}
		params := listMailboxActionParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListMailboxActionLogsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListMailboxActionCmd, utils.StringArg("order-id"))
	ListMailboxActionCmd.Flags().StringP("email", "", "", "Mailbox email address. Must belong to the order's domain.")
	ListMailboxActionCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListMailboxActionCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListMailboxActionCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListMailboxActionCmd.Flags().StringP("event", "", "", "Filter mailbox action log entries by event type (one of: MessageNew, MessageRead, MessageAppend, MessageExpunge, MailboxCreate, MailboxDelete, MailboxRename)")
	ListMailboxActionCmd.Flags().IntP("page", "", 0, "Page number")
	ListMailboxActionCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListMailboxActionCmd.MarkFlagRequired("email")
	ListMailboxActionCmd.RegisterFlagCompletionFunc("email", cobra.NoFileCompletions)
	ListMailboxActionCmd.RegisterFlagCompletionFunc("date", cobra.NoFileCompletions)
	ListMailboxActionCmd.RegisterFlagCompletionFunc("from-date", cobra.NoFileCompletions)
	ListMailboxActionCmd.RegisterFlagCompletionFunc("to-date", cobra.NoFileCompletions)
	ListMailboxActionCmd.RegisterFlagCompletionFunc("event", cobra.FixedCompletions([]cobra.Completion{"MessageNew", "MessageRead", "MessageAppend", "MessageExpunge", "MailboxCreate", "MailboxDelete", "MailboxRename"}, cobra.ShellCompDirectiveNoFileComp))
	ListMailboxActionCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListMailboxActionCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListMailboxActionCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListMailboxActionCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listMailboxActionParams(cmd *cobra.Command) *client.MailListMailboxActionLogsV1Params {
	params := &client.MailListMailboxActionLogsV1Params{}
	emailVal, _ := cmd.Flags().GetString("email")
	params.Email = utils.StringToEmail(emailVal)
	if cmd.Flags().Changed("date") {
		v, _ := cmd.Flags().GetString("date")
		t := utils.StringToDate(v)
		params.Date = &t
	}
	if cmd.Flags().Changed("from-date") {
		v, _ := cmd.Flags().GetString("from-date")
		t := utils.StringToTime(v)
		params.FromDate = &t
	}
	if cmd.Flags().Changed("to-date") {
		v, _ := cmd.Flags().GetString("to-date")
		t := utils.StringToTime(v)
		params.ToDate = &t
	}
	if cmd.Flags().Changed("event") {
		v, _ := cmd.Flags().GetString("event")
		e := client.MailListMailboxActionLogsV1ParamsEvent(v)
		params.Event = &e
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
