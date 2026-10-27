package logs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListAccessCmd = &cobra.Command{
	Use:     "list-access [order-id]",
	Short:   "List access logs",
	Long:    "Retrieve paginated access logs for the domain attached to the given\nmail order. Supports filtering by account, date range, protocol,\nstatus, and deletion flag. Results are sorted by timestamp descending.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail logs list-access OR1a2b3c4d5e6f7g",
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
		if err := utils.EnumCheck(cmd, "status", []string{"Successful", "Failed"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "protocol", []string{"imap", "pop3", "smtp"}); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "account", "email"); err != nil {
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
		params := listAccessParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListAccessLogsV1WithResponse(ctx, orderID, params)
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
	utils.RegisterScopedArgs(ListAccessCmd, 1, utils.StringArg("order-id"))
	ListAccessCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListAccessCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListAccessCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListAccessCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListAccessCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListAccessCmd.Flags().StringP("protocol", "", "", "Filter access log entries by protocol (one of: imap, pop3, smtp)")
	ListAccessCmd.Flags().BoolP("has-deletions", "", false, "Filter access log entries by whether the session had deletions")
	ListAccessCmd.Flags().IntP("page", "", 0, "Page number")
	ListAccessCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListAccessCmd.RegisterFlagCompletionFunc("account", cobra.NoFileCompletions)
	ListAccessCmd.RegisterFlagCompletionFunc("date", cobra.NoFileCompletions)
	ListAccessCmd.RegisterFlagCompletionFunc("from-date", cobra.NoFileCompletions)
	ListAccessCmd.RegisterFlagCompletionFunc("to-date", cobra.NoFileCompletions)
	ListAccessCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"Successful", "Failed"}, cobra.ShellCompDirectiveNoFileComp))
	ListAccessCmd.RegisterFlagCompletionFunc("protocol", cobra.FixedCompletions([]cobra.Completion{"imap", "pop3", "smtp"}, cobra.ShellCompDirectiveNoFileComp))
	ListAccessCmd.RegisterFlagCompletionFunc("has-deletions", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ListAccessCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListAccessCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListAccessCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListAccessCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listAccessParams(cmd *cobra.Command) *client.MailListAccessLogsV1Params {
	params := &client.MailListAccessLogsV1Params{}
	if cmd.Flags().Changed("account") {
		v, _ := cmd.Flags().GetString("account")
		t := utils.StringToEmail(v)
		params.Account = &t
	}
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
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.MailListAccessLogsV1ParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("protocol") {
		v, _ := cmd.Flags().GetString("protocol")
		e := client.MailListAccessLogsV1ParamsProtocol(v)
		params.Protocol = &e
	}
	if cmd.Flags().Changed("has-deletions") {
		v, _ := cmd.Flags().GetBool("has-deletions")
		params.HasDeletions = &v
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
