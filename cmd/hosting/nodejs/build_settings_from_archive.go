package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildSettingsFromArchiveCmd = &cobra.Command{
	Use:     "build-settings-from-archive <username> <domain>",
	Short:   "Get Node.js build settings from archive",
	Long:    "Auto-detect Node.js build settings from a package.json inside an archive already on the server.\n\nUse this before calling `Start Node.js Build` to preview what settings will be used,\nor to let the user review and override values (framework, node version, root directory,\noutput directory, build script) before committing to a build.\n\nThe archive must already be present on the website's file storage. Use the\n`Generate Upload URL` endpoint to obtain credentials and upload the archive first.",
	Example: "  hostinger hosting nodejs build-settings-from-archive u123456789 mydomain.tld --archive-path value",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := buildSettingsFromArchiveParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsBuildSettingsFromArchiveV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(BuildSettingsFromArchiveCmd, utils.StringArg("username"), utils.StringArg("domain"))
	BuildSettingsFromArchiveCmd.Flags().StringP("archive-path", "", "", "The path to the archive file relative to the document root of the vhost")
	BuildSettingsFromArchiveCmd.MarkFlagRequired("archive-path")
	BuildSettingsFromArchiveCmd.RegisterFlagCompletionFunc("archive-path", cobra.NoFileCompletions)
}

func buildSettingsFromArchiveParams(cmd *cobra.Command) *client.HostingGetNodeJsBuildSettingsFromArchiveV1Params {
	params := &client.HostingGetNodeJsBuildSettingsFromArchiveV1Params{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	params.ArchivePath = archivePathVal
	return params
}
