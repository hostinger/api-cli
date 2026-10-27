package nodejs

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:     "build [domain] [build-uuid]",
	Short:   "Get Node.js build details",
	Long:    "Returns one build by UUID: its state (`pending`, `running`, `completed`, `failed`), the\noptions it ran with and timestamps. Poll this while a build is pending or running. When it\nis failed, read `Get NodeJS build logs` and `Analyse failed Node.js build` for the cause.\nReturns 404 when the UUID does not belong to a build of this website.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.\n\nThe build can be given as [build-uuid] or left out, in which case it is the build `hostinger up` last started for the website (remembered in .hostinger/site.json), or a choice among the website's builds when run in a terminal.",
	Example: "  hostinger hosting nodejs build mydomain.tld 123e4567-e89b-12d3-a456-426614174000",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsBuildDetailsV1WithResponse(ctx, username, domain, uuid.MustParse(buildUUID))
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
	utils.RegisterScopedArgs(BuildCmd, 3, utils.StringArg("username"), utils.StringArg("domain"), utils.StringArg("uuid"))
}
