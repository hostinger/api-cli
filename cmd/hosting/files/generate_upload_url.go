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

var GenerateUploadUrlCmd = &cobra.Command{
	Use:   "generate-upload-url",
	Short: "Generate upload URL",
	Long:  "Generate a file browser upload URL with authentication credentials\nfor uploading files directly to a website's file storage.\n\nReturns `url`, `auth_key` and `rest_auth_key`. Use these to upload a file to the\nwebsite's `public_html` directory via the TUS resumable upload protocol (TUS 1.0.0).\nSend `X-Auth: {auth_key}` and `X-Auth-Rest: {rest_auth_key}` headers on every request\nbelow.\n\n1. Create the upload: `POST` to `{url}/{relative_file_path}?override=true` with headers\n   `upload-length: {file size in bytes}` and `upload-offset: 0`. Expect `201 Created`.\n2. Upload the file: send the file bytes to the same location (any TUS 1.0.0 client, or\n   `PATCH` requests with an `upload-offset` header tracking progress) until complete.\n\n`relative_file_path` is the destination path inside `public_html`, e.g. `app.zip`.\n\nInstead of a TUS client, plain `curl` also works:\n```\nFILE=app.zip\nSIZE=$(stat -f%z \"$FILE\")   # stat -c%s on Linux\n\ncurl -i -X POST \"{url}/${FILE}?override=true\" \\\n  -H \"X-Auth: {auth_key}\" \\\n  -H \"X-Auth-Rest: {rest_auth_key}\" \\\n  -H \"Tus-Resumable: 1.0.0\" \\\n  -H \"Upload-Length: ${SIZE}\" \\\n  -H \"Upload-Offset: 0\"\n# -> 201 Created\n\ncurl -i -X PATCH \"{url}/${FILE}?override=true\" \\\n  -H \"X-Auth: {auth_key}\" \\\n  -H \"X-Auth-Rest: {rest_auth_key}\" \\\n  -H \"Tus-Resumable: 1.0.0\" \\\n  -H \"Content-Type: application/offset+octet-stream\" \\\n  -H \"Upload-Offset: 0\" \\\n  --data-binary \"@${FILE}\"\n# -> 204 No Content, Upload-Offset response header equals SIZE when done\n```",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(generateUploadUrlBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingGenerateUploadURLV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	GenerateUploadUrlCmd.Flags().StringP("domain", "", "", "Website domain")
	GenerateUploadUrlCmd.Flags().StringP("username", "", "", "Account username")
	GenerateUploadUrlCmd.MarkFlagRequired("domain")
	GenerateUploadUrlCmd.MarkFlagRequired("username")
}

func generateUploadUrlBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	usernameVal, _ := cmd.Flags().GetString("username")
	body["username"] = usernameVal
	return body
}
