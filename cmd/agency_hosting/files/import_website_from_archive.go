package files

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ImportWebsiteFromArchiveCmd = &cobra.Command{
	Use:               "import-website-from-archive <website_uid>",
	Short:             "Import website from archive",
	Long:              "Imports an Agency Plan website from an already-uploaded archive.\n\nUpload the archive to the website's root directory via file browser first, then provide its\nfilename in this request. Website contents are overwritten by the archive contents. Supported\narchive types: .zip, .tar, .tar.gz, .tgz.",
	Example:           "  hostinger agency-hosting files import-website-from-archive zpwlGlp19 --archive-name archive.zip",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := importWebsiteFromArchiveBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingImportWebsiteFromArchiveV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ImportWebsiteFromArchiveCmd, utils.StringArg("website_uid"))
	ImportWebsiteFromArchiveCmd.Flags().StringP("archive-name", "", "", "Archive filename (e.g., archive.zip). The file must already be uploaded to the website's .h5g/ directory.")
	ImportWebsiteFromArchiveCmd.MarkFlagRequired("archive-name")
	ImportWebsiteFromArchiveCmd.RegisterFlagCompletionFunc("archive-name", cobra.NoFileCompletions)
}

func importWebsiteFromArchiveBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	archiveNameVal, _ := cmd.Flags().GetString("archive-name")
	body["archive_name"] = archiveNameVal
	return body, nil
}
