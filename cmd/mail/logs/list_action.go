package logs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListActionCmd = &cobra.Command{
	Use:     "list-action <order-id>",
	Short:   "List action logs",
	Long:    "Retrieve paginated account action logs (administrative and user\nactions) for the given mail order. Supports filtering by account,\ndate range, and status. Results are sorted by timestamp descending.",
	Example: "  hostinger mail logs list-action OR1a2b3c4d5e6f7g",
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
		params := listActionParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListActionLogsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListActionCmd, utils.StringArg("order-id"))
	ListActionCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListActionCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListActionCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListActionCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListActionCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListActionCmd.Flags().IntP("page", "", 0, "Page number")
	ListActionCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListActionCmd.RegisterFlagCompletionFunc("account", cobra.NoFileCompletions)
	ListActionCmd.RegisterFlagCompletionFunc("date", cobra.NoFileCompletions)
	ListActionCmd.RegisterFlagCompletionFunc("from-date", cobra.NoFileCompletions)
	ListActionCmd.RegisterFlagCompletionFunc("to-date", cobra.NoFileCompletions)
	ListActionCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"Successful", "Failed"}, cobra.ShellCompDirectiveNoFileComp))
	ListActionCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListActionCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListActionCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListActionCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listActionParams(cmd *cobra.Command) *client.MailListActionLogsV1Params {
	params := &client.MailListActionLogsV1Params{}
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
		e := client.MailListActionLogsV1ParamsStatus(v)
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
