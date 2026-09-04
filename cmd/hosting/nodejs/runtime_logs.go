package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RuntimeLogsCmd = &cobra.Command{
	Use:   "runtime-logs <username> <domain>",
	Short: "Get Node.js runtime logs",
	Long:  "Returns the Node.js application's runtime console log entries, oldest first, each with\ntimestamp, level and message. On the first call send `period` (`1h`, `1d`, `1w` or `1m`)\nand optionally `levels` and `limit` (1-5000, default 1000); when more entries match than\n`limit`, the newest are kept.\n\nTo poll for new entries send `total_lines + 1` from the previous response as `from_line`\nand omit `period`; `period` and `from_line` cannot be combined. Lines that are not JSON\nwith a timestamp, level and message are skipped, so `logs` may hold fewer than `limit`\nentries while `total_lines` counts every raw line. Entries with a timestamp before\n`last_deployed_at` belong to the previous deployment. Returns an empty `logs` list when\nthe application has not written a log file yet.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "period", []string{"1h", "1d", "1w", "1m"})
		r, err := api.Request().HostingGetNodeJsRuntimeLogsV1WithResponse(context.TODO(), args[0], args[1], runtimeLogsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RuntimeLogsCmd.Flags().StringP("period", "", "", "Time window for the first fetch. Required when `from_line` is not sent. (one of: 1h, 1d, 1w, 1m)")
	RuntimeLogsCmd.Flags().IntP("from-line", "", 0, "1-based line of the log file to start from. For polling send `total_lines + 1` from the\nprevious response. Cannot be combined with `period`.")
	RuntimeLogsCmd.Flags().IntP("limit", "", 1000, "Maximum number of log entries to return. When more entries match, the newest are kept.")
	RuntimeLogsCmd.Flags().StringSliceP("levels", "", nil, "Return only entries with these log levels, sent as a comma-separated list, e.g. ERROR,WARN.\nMatching runs on the raw log line, so entries written with numeric levels (for example by\npino) are excluded while this filter is set. (one of: LOG, ERROR, WARN, INFO, DEBUG, TRACE)")
}

func runtimeLogsParams(cmd *cobra.Command) *client.HostingGetNodeJsRuntimeLogsV1Params {
	params := &client.HostingGetNodeJsRuntimeLogsV1Params{}
	if cmd.Flags().Changed("period") {
		v, _ := cmd.Flags().GetString("period")
		e := client.HostingGetNodeJsRuntimeLogsV1ParamsPeriod(v)
		params.Period = &e
	}
	if cmd.Flags().Changed("from-line") {
		v, _ := cmd.Flags().GetInt("from-line")
		params.FromLine = &v
	}
	if cmd.Flags().Changed("limit") {
		v, _ := cmd.Flags().GetInt("limit")
		params.Limit = &v
	}
	if cmd.Flags().Changed("levels") {
		v, _ := cmd.Flags().GetStringSlice("levels")
		es := make([]client.HostingGetNodeJsRuntimeLogsV1ParamsLevels, len(v))
		for i, s := range v {
			es[i] = client.HostingGetNodeJsRuntimeLogsV1ParamsLevels(s)
		}
		params.Levels = &es
	}
	return params
}
