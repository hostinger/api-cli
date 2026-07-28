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

var ListInboundCmd = &cobra.Command{
	Use:   "list-inbound <order-id>",
	Short: "List inbound logs",
	Long:  "Retrieve paginated inbound (received mail) delivery logs for the\ndomain attached to the given mail order. Supports filtering by\naccount, date range, status, sender, and recipient. Results are\nsorted by timestamp descending.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"Successful", "Failed"})
		r, err := api.Request().MailListInboundLogsV1WithResponse(context.TODO(), args[0], listInboundParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListInboundCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListInboundCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListInboundCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListInboundCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListInboundCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListInboundCmd.Flags().StringP("sender", "", "", "Filter log entries by sender. Accepts a full email address or a domain.")
	ListInboundCmd.Flags().StringP("recipient", "", "", "Filter log entries by recipient. Accepts a full email address or a domain.")
	ListInboundCmd.Flags().IntP("page", "", 0, "Page number")
	ListInboundCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listInboundParams(cmd *cobra.Command) *client.MailListInboundLogsV1Params {
	params := &client.MailListInboundLogsV1Params{}
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
		e := client.MailListInboundLogsV1ParamsStatus(v)
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
