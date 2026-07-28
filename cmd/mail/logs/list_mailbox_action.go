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

var ListMailboxActionCmd = &cobra.Command{
	Use:   "list-mailbox-action <order-id>",
	Short: "List mailbox action logs",
	Long:  "Retrieve paginated mailbox action logs (message and mailbox events)\nfor a mailbox in the given mail order. The mailbox email must belong\nto the order's domain. Supports date range and event type filters.\nResults are sorted by timestamp descending.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "event", []string{"MessageNew", "MessageRead", "MessageAppend", "MessageExpunge", "MailboxCreate", "MailboxDelete", "MailboxRename"})
		r, err := api.Request().MailListMailboxActionLogsV1WithResponse(context.TODO(), args[0], listMailboxActionParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListMailboxActionCmd.Flags().StringP("email", "", "", "Mailbox email address. Must belong to the order's domain.")
	ListMailboxActionCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListMailboxActionCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListMailboxActionCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListMailboxActionCmd.Flags().StringP("event", "", "", "Filter mailbox action log entries by event type (one of: MessageNew, MessageRead, MessageAppend, MessageExpunge, MailboxCreate, MailboxDelete, MailboxRename)")
	ListMailboxActionCmd.Flags().IntP("page", "", 0, "Page number")
	ListMailboxActionCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListMailboxActionCmd.MarkFlagRequired("email")
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
