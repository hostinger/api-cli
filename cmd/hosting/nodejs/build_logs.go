package nodejs

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var BuildLogsCmd = &cobra.Command{
	Use:   "build-logs <username> <domain> <uuid>",
	Short: "Get NodeJS build logs",
	Long:  "Retrieve logs from a specific Node.js build process.\n\nTo stream live output while a build is running, poll this endpoint repeatedly\nwhile the build state is `running`, passing the previously returned `lines` count\nas `from_line` to fetch only new output since the last call.\nLog content may contain ANSI escape sequences (color codes).",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetNodeJSBuildLogsV1WithResponse(context.TODO(), args[0], args[1], uuid.MustParse(args[2]), buildLogsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	BuildLogsCmd.Flags().IntP("from-line", "", 0, "Line from which to start retrieving logs")
}

func buildLogsParams(cmd *cobra.Command) *client.HostingGetNodeJSBuildLogsV1Params {
	params := &client.HostingGetNodeJSBuildLogsV1Params{}
	if cmd.Flags().Changed("from-line") {
		v, _ := cmd.Flags().GetInt("from-line")
		params.FromLine = &v
	}
	return params
}
