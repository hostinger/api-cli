package nodejs

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildLogsCmd = &cobra.Command{
	Use:     "build-logs <username> <domain> <uuid>",
	Short:   "Get NodeJS build logs",
	Long:    "Retrieve logs from a specific Node.js build process.\n\nTo stream live output while a build is running, poll this endpoint repeatedly\nwhile the build state is `running`, passing the previously returned `lines` count\nas `from_line` to fetch only new output since the last call.\nLog content may contain ANSI escape sequences (color codes).",
	Example: "  hostinger hosting nodejs build-logs u123456789 mydomain.tld 123e4567-e89b-12d3-a456-426614174000",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := buildLogsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJSBuildLogsV1WithResponse(ctx, args[0], args[1], uuid.MustParse(args[2]), params)
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
	utils.RegisterArgs(BuildLogsCmd, utils.StringArg("username"), utils.StringArg("domain"), utils.UUIDArg("uuid"))
	BuildLogsCmd.Flags().IntP("from-line", "", 0, "Line from which to start retrieving logs")
	BuildLogsCmd.RegisterFlagCompletionFunc("from-line", cobra.NoFileCompletions)
}

func buildLogsParams(cmd *cobra.Command) *client.HostingGetNodeJSBuildLogsV1Params {
	params := &client.HostingGetNodeJSBuildLogsV1Params{}
	if cmd.Flags().Changed("from-line") {
		v, _ := cmd.Flags().GetInt("from-line")
		params.FromLine = &v
	}
	return params
}
