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

var ListOutboundCmd = &cobra.Command{
	Use:   "list-outbound <order-id>",
	Short: "List outbound logs",
	Long:  "Retrieve paginated outbound (sent mail) delivery logs for the domain\nattached to the given mail order. Supports filtering by account, date\nrange, status, sender, and recipient. Results are sorted by timestamp\ndescending.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"Successful", "Failed"})
		r, err := api.Request().MailListOutboundLogsV1WithResponse(context.TODO(), args[0], listOutboundParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListOutboundCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListOutboundCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListOutboundCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListOutboundCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListOutboundCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListOutboundCmd.Flags().StringP("sender", "", "", "Filter log entries by sender. Accepts a full email address or a domain.")
	ListOutboundCmd.Flags().StringP("recipient", "", "", "Filter log entries by recipient. Accepts a full email address or a domain.")
	ListOutboundCmd.Flags().IntP("page", "", 0, "Page number")
	ListOutboundCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
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
