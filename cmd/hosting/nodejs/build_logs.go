package nodejs

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildLogsCmd = &cobra.Command{
	Use:     "build-logs [domain] [build-uuid]",
	Short:   "Get NodeJS build logs",
	Long:    "Retrieve logs from a specific Node.js build process.\n\nTo stream live output while a build is running, poll this endpoint repeatedly\nwhile the build state is `running`, passing the previously returned `lines` count\nas `from_line` to fetch only new output since the last call.\nLog content may contain ANSI escape sequences (color codes).\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.\n\nThe build can be given as [build-uuid] or left out, in which case it is the build `hostinger up` last started for the website (remembered in .hostinger/site.json), or a choice among the website's builds when run in a terminal.",
	Example: "  hostinger hosting nodejs build-logs mydomain.tld 123e4567-e89b-12d3-a456-426614174000",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, domain, buildUUID, args, err := site.ResolveBuild(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		params := buildLogsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJSBuildLogsV1WithResponse(ctx, username, domain, uuid.MustParse(buildUUID), params)
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
	utils.RegisterScopedArgs(BuildLogsCmd, 3, utils.StringArg("username"), utils.StringArg("domain"), utils.StringArg("uuid"))
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
