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

var DeployStaticSiteArchiveCmd = &cobra.Command{
	Use:   "deploy-static-site-archive <username> <domain>",
	Short: "Deploy static site archive",
	Long:  "Deploy a static application from an archive file.\n\nWARNING: this overwrites the website's existing contents and cannot be undone —\nverify this is intended before calling this endpoint.\n\nThis endpoint allows you to deploy a static application from an archive\nfile that has been uploaded to the website's directory.\n\nThis only works for static sites (pre-built HTML/CSS/JS with no build step). For\nNode.js applications, use `Create NodeJS build from archive` instead, or\n`Start Node.js build` if the archive is already uploaded. For WordPress sites,\nuse `Import WordPress website`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deployStaticSiteArchiveBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingDeployStaticSiteArchiveV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeployStaticSiteArchiveCmd.Flags().StringP("archive-path", "", "", "Relative path to the archive file from website root directory")
	DeployStaticSiteArchiveCmd.MarkFlagRequired("archive-path")
}

func deployStaticSiteArchiveBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	return body
}
