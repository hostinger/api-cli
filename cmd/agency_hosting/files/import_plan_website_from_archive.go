package files

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ImportPlanWebsiteFromArchiveCmd = &cobra.Command{
	Use:   "import-plan-website-from-archive <website_uid>",
	Short: "Import Agency Plan website from archive",
	Long:  "Imports an Agency Plan website from an already-uploaded archive.\n\nUpload the archive to the website's root directory via file browser first, then provide its\nfilename in this request. Website contents are overwritten by the archive contents. Supported\narchive types: .zip, .tar, .tar.gz, .tgz.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(importPlanWebsiteFromArchiveBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingImportAgencyPlanWebsiteFromArchiveV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ImportPlanWebsiteFromArchiveCmd.Flags().StringP("archive-name", "", "", "Archive filename (e.g., archive.zip). The file must already be uploaded to the website's .h5g/ directory.")
	ImportPlanWebsiteFromArchiveCmd.MarkFlagRequired("archive-name")
}

func importPlanWebsiteFromArchiveBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	archiveNameVal, _ := cmd.Flags().GetString("archive-name")
	body["archive_name"] = archiveNameVal
	return body
}
