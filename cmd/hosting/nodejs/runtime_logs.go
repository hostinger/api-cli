package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RuntimeLogsCmd = &cobra.Command{
	Use:     "runtime-logs [domain]",
	Short:   "Get Node.js runtime logs",
	Long:    "Returns the Node.js application's runtime console log entries, oldest first, each with\ntimestamp, level and message. On the first call send `period` (`1h`, `1d`, `1w` or `1m`)\nand optionally `levels` and `limit` (1-5000, default 1000); when more entries match than\n`limit`, the newest are kept.\n\nTo poll for new entries send `total_lines + 1` from the previous response as `from_line`\nand omit `period`; `period` and `from_line` cannot be combined. Lines that are not JSON\nwith a timestamp, level and message are skipped, so `logs` may hold fewer than `limit`\nentries while `total_lines` counts every raw line. Entries with a timestamp before\n`last_deployed_at` belong to the previous deployment. Returns an empty `logs` list when\nthe application has not written a log file yet.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting nodejs runtime-logs mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		if err := utils.EnumCheck(cmd, "period", []string{"1h", "1d", "1w", "1m"}); err != nil {
			return err
		}
		params := runtimeLogsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsRuntimeLogsV1WithResponse(ctx, username, domain, params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterScopedArgs(RuntimeLogsCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	RuntimeLogsCmd.Flags().StringP("period", "", "", "Time window for the first fetch. Required when `from_line` is not sent. (one of: 1h, 1d, 1w, 1m)")
	RuntimeLogsCmd.Flags().IntP("from-line", "", 0, "1-based line of the log file to start from. For polling send `total_lines + 1` from the previous response. Cannot be combined with `period`.")
	RuntimeLogsCmd.Flags().IntP("limit", "", 1000, "Maximum number of log entries to return. When more entries match, the newest are kept.")
	RuntimeLogsCmd.Flags().StringSliceP("levels", "", nil, "Return only entries with these log levels, sent as a comma-separated list, e.g. ERROR,WARN. Matching runs on the raw log line, so entries written with numeric levels (for example by pino) are excluded while this filter is set. (one of: LOG, ERROR, WARN, INFO, DEBUG, TRACE)")
	RuntimeLogsCmd.RegisterFlagCompletionFunc("period", cobra.FixedCompletions([]cobra.Completion{"1h", "1d", "1w", "1m"}, cobra.ShellCompDirectiveNoFileComp))
	RuntimeLogsCmd.RegisterFlagCompletionFunc("from-line", cobra.NoFileCompletions)
	RuntimeLogsCmd.RegisterFlagCompletionFunc("limit", cobra.NoFileCompletions)
	RuntimeLogsCmd.RegisterFlagCompletionFunc("levels", cobra.FixedCompletions([]cobra.Completion{"LOG", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}, cobra.ShellCompDirectiveNoFileComp))
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
