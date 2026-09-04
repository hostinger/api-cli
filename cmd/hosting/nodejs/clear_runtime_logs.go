package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ClearRuntimeLogsCmd = &cobra.Command{
	Use:   "clear-runtime-logs <username> <domain>",
	Short: "Clear Node.js runtime logs",
	Long:  "Empties the Node.js application's runtime log file. This cannot be undone, so confirm with\nthe user before calling it. Returns success even when no log file exists yet.\n\nUse it before reproducing a problem so the next `Get Node.js runtime logs` call returns\nonly fresh entries; start that call with `period` again instead of reusing a `from_line`\nfrom before the clear.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingClearNodeJsRuntimeLogsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
