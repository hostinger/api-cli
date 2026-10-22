package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildSettingsCmd = &cobra.Command{
	Use:     "build-settings <username> <domain>",
	Short:   "Get Node.js build settings",
	Long:    "Returns the build settings stored for the website: framework (`app_type`), Node.js version,\nroot and output directory, build script, entry file and package manager. Stored settings\ndrive Git auto-deployment builds. A build started through the API uses the values sent in\nthat request and saves them here only when no settings exist yet.\n\nReturns 404 until the first build or the first settings update stores them. Use this after\na failed build to check whether the framework or the entry file were detected wrong, then\nfix them with the `Update Node.js build settings` endpoint.",
	Example: "  hostinger hosting nodejs build-settings u123456789 mydomain.tld",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsBuildSettingsV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(BuildSettingsCmd, utils.StringArg("username"), utils.StringArg("domain"))
}
