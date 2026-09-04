package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var BuildSettingsCmd = &cobra.Command{
	Use:   "build-settings <username> <domain>",
	Short: "Get Node.js build settings",
	Long:  "Returns the build settings stored for the website: framework (`app_type`), Node.js version,\nroot and output directory, build script, entry file and package manager. Stored settings\ndrive Git auto-deployment builds. A build started through the API uses the values sent in\nthat request and saves them here only when no settings exist yet.\n\nReturns 404 until the first build or the first settings update stores them. Use this after\na failed build to check whether the framework or the entry file were detected wrong, then\nfix them with the `Update Node.js build settings` endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetNodeJsBuildSettingsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
