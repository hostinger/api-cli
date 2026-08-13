package websites

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var BuildNodejsAssetsCmd = &cobra.Command{
	Use:   "build-nodejs-assets <website_uid>",
	Short: "Build website NodeJS assets",
	Long:  "Builds and deploys a Node.js application for an Agency Plan website from an already-uploaded archive.\n\nUpload the archive to file browser first, then provide its relative path from document root in this request.\nWebsite contents are overwritten by the build result, which is deployed to public_html.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(buildNodejsAssetsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingBuildWebsiteNodeJSAssetsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	BuildNodejsAssetsCmd.Flags().StringP("archive-path", "", "", "Directory, relative to the website document root, where the uploaded site archive currently lives. Most commonly this is simply `public_html`.")
	BuildNodejsAssetsCmd.MarkFlagRequired("archive-path")
}

func buildNodejsAssetsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	return body
}
