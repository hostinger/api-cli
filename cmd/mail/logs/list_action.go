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

var ListActionCmd = &cobra.Command{
	Use:   "list-action <order-id>",
	Short: "List action logs",
	Long:  "Retrieve paginated account action logs (administrative and user\nactions) for the given mail order. Supports filtering by account,\ndate range, and status. Results are sorted by timestamp descending.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"Successful", "Failed"})
		r, err := api.Request().MailListActionLogsV1WithResponse(context.TODO(), args[0], listActionParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListActionCmd.Flags().StringP("account", "", "", "Filter log entries by a specific email account")
	ListActionCmd.Flags().StringP("date", "", "", "Exact date filter (YYYY-MM-DD). Takes precedence over `from_date`/`to_date` when both are given.")
	ListActionCmd.Flags().StringP("from-date", "", "", "Date range start (RFC 3339)")
	ListActionCmd.Flags().StringP("to-date", "", "", "Date range end (RFC 3339)")
	ListActionCmd.Flags().StringP("status", "", "", "Filter log entries by status (one of: Successful, Failed)")
	ListActionCmd.Flags().IntP("page", "", 0, "Page number")
	ListActionCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
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
