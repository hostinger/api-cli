package nodejs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var BuildSettingsFromArchiveCmd = &cobra.Command{
	Use:   "build-settings-from-archive <username> <domain>",
	Short: "Get Node.js build settings from archive",
	Long:  "Auto-detect Node.js build settings from a package.json inside an archive already on the server.\n\nUse this before calling `Start Node.js Build` to preview what settings will be used,\nor to let the user review and override values (framework, node version, root directory,\noutput directory, build script) before committing to a build.\n\nThe archive must already be present on the website's file storage. Use the\n`Generate Upload URL` endpoint to obtain credentials and upload the archive first.\nTo upload an archive and start a build in one step without inspecting settings first,\nuse the `Create Node.js Build from Archive` endpoint instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetNodeJsBuildSettingsFromArchiveV1WithResponse(context.TODO(), args[0], args[1], buildSettingsFromArchiveParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	BuildSettingsFromArchiveCmd.Flags().StringP("archive-path", "", "", "The path to the archive file relative to the document root of the vhost")
	BuildSettingsFromArchiveCmd.MarkFlagRequired("archive-path")
}

func buildSettingsFromArchiveParams(cmd *cobra.Command) *client.HostingGetNodeJsBuildSettingsFromArchiveV1Params {
	params := &client.HostingGetNodeJsBuildSettingsFromArchiveV1Params{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	params.ArchivePath = archivePathVal
	return params
}
