package logs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListAccessCmd = &cobra.Command{
	Use:   "list-access <order-id>",
	Short: "List access logs",
	Long:  "Retrieve paginated access logs for the domain attached to the given\nmail order. Supports filtering by account, date range, protocol,\nstatus, and deletion flag. Results are sorted by timestamp descending.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"Successful", "Failed"})
		utils.EnumCheck(cmd, "protocol", []string{"imap", "pop3", "smtp"})
		r, err := api.Request().MailListAccessLogsV1WithResponse(context.TODO(), args[0], listAccessParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListAccessCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListAccessCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListAccessCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListAccessCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListAccessCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListAccessCmd.Flags().StringP("protocol", "", "", "Filter access log entries by protocol (one of: imap, pop3, smtp)")
	ListAccessCmd.Flags().BoolP("has-deletions", "", false, "Filter access log entries by whether the session had deletions")
	ListAccessCmd.Flags().IntP("page", "", 0, "Page number")
	ListAccessCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
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
