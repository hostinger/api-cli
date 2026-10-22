package files

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GenerateUploadUrlCmd = &cobra.Command{
	Use:               "generate-upload-url <website_uid>",
	Short:             "Generate upload URL",
	Long:              "Generate a file browser upload URL with authentication credentials for uploading files\nto an Agency Plan website's file storage.\n\nReturns `url`, `auth_key` and `rest_auth_key`. Use these to upload a file to the\nwebsite's file storage via the TUS resumable upload protocol (TUS 1.0.0). Send\n`X-Auth: {auth_key}` and `X-Auth-Rest: {rest_auth_key}` headers on every request below.\n\n1. Create the upload: `POST` to `{url}/{relative_file_path}?override=true` with headers\n   `upload-length: {file size in bytes}` and `upload-offset: 0`. Expect `201 Created`.\n2. Upload the file: send the file bytes to the same location (any TUS 1.0.0 client, or\n   `PATCH` requests with an `upload-offset` header tracking progress) until complete.\n\n`relative_file_path` is the destination path inside the website's file storage, e.g.\n`app.zip`.\n\nInstead of a TUS client, plain `curl` also works:\n```\nFILE=app.zip\nSIZE=$(stat -f%z \"$FILE\")   # stat -c%s on Linux\n\ncurl -i -X POST \"{url}/${FILE}?override=true\" \\\n  -H \"X-Auth: {auth_key}\" \\\n  -H \"X-Auth-Rest: {rest_auth_key}\" \\\n  -H \"Tus-Resumable: 1.0.0\" \\\n  -H \"Upload-Length: ${SIZE}\" \\\n  -H \"Upload-Offset: 0\"\n-> 201 Created\n\ncurl -i -X PATCH \"{url}/${FILE}?override=true\" \\\n  -H \"X-Auth: {auth_key}\" \\\n  -H \"X-Auth-Rest: {rest_auth_key}\" \\\n  -H \"Tus-Resumable: 1.0.0\" \\\n  -H \"Content-Type: application/offset+octet-stream\" \\\n  -H \"Upload-Offset: 0\" \\\n  --data-binary \"@${FILE}\"\n-> 204 No Content, Upload-Offset response header equals SIZE when done\n```",
	Example:           "  hostinger agency-hosting files generate-upload-url zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingGenerateUploadURLV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(GenerateUploadUrlCmd, utils.StringArg("website_uid"))
}
