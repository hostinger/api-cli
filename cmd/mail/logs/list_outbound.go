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

var ListOutboundCmd = &cobra.Command{
	Use:     "list-outbound [order-id]",
	Short:   "List outbound logs",
	Long:    "Retrieve paginated outbound (sent mail) delivery logs for the domain\nattached to the given mail order. Supports filtering by account, date\nrange, status, sender, and recipient. Results are sorted by timestamp\ndescending.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail logs list-outbound OR1a2b3c4d5e6f7g",
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
		params := listOutboundParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListOutboundLogsV1WithResponse(ctx, orderID, params)
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
	utils.RegisterScopedArgs(ListOutboundCmd, 1, utils.StringArg("order-id"))
	ListOutboundCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListOutboundCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListOutboundCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListOutboundCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListOutboundCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListOutboundCmd.Flags().StringP("sender", "", "", "Filter log entries by sender. Accepts a full email address or a domain.")
	ListOutboundCmd.Flags().StringP("recipient", "", "", "Filter log entries by recipient. Accepts a full email address or a domain.")
	ListOutboundCmd.Flags().IntP("page", "", 0, "Page number")
	ListOutboundCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListOutboundCmd.RegisterFlagCompletionFunc("account", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("date", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("from-date", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("to-date", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"Successful", "Failed"}, cobra.ShellCompDirectiveNoFileComp))
	ListOutboundCmd.RegisterFlagCompletionFunc("sender", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("recipient", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListOutboundCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListOutboundCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListOutboundCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listOutboundParams(cmd *cobra.Command) *client.MailListOutboundLogsV1Params {
	params := &client.MailListOutboundLogsV1Params{}
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
		e := client.MailListOutboundLogsV1ParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("sender") {
		v, _ := cmd.Flags().GetString("sender")
		params.Sender = &v
	}
	if cmd.Flags().Changed("recipient") {
		v, _ := cmd.Flags().GetString("recipient")
		params.Recipient = &v
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
