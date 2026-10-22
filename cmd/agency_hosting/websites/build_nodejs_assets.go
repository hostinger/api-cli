package websites

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildNodejsAssetsCmd = &cobra.Command{
	Use:               "build-nodejs-assets <website_uid>",
	Short:             "Build website NodeJS assets",
	Long:              "Builds and deploys a Node.js application for an Agency Plan website from an already-uploaded archive.\n\nUpload the archive to file browser first, then provide its relative path from document root in this request.\nWebsite contents are overwritten by the build result, which is deployed to public_html.",
	Example:           "  hostinger agency-hosting websites build-nodejs-assets zpwlGlp19 --archive-path public_html",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := buildNodejsAssetsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingBuildWebsiteNodeJSAssetsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(BuildNodejsAssetsCmd, utils.StringArg("website_uid"))
	BuildNodejsAssetsCmd.Flags().StringP("archive-path", "", "", "Directory, relative to the website document root, where the uploaded site archive currently lives. Most commonly this is simply `public_html`.")
	BuildNodejsAssetsCmd.MarkFlagRequired("archive-path")
	BuildNodejsAssetsCmd.RegisterFlagCompletionFunc("archive-path", cobra.NoFileCompletions)
}

func buildNodejsAssetsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	return body, nil
}
