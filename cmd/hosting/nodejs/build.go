package nodejs

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:     "build <username> <domain> <uuid>",
	Short:   "Get Node.js build details",
	Long:    "Returns one build by UUID: its state (`pending`, `running`, `completed`, `failed`), the\noptions it ran with and timestamps. Poll this while a build is pending or running. When it\nis failed, read `Get NodeJS build logs` and `Analyse failed Node.js build` for the cause.\nReturns 404 when the UUID does not belong to a build of this website.",
	Example: "  hostinger hosting nodejs build u123456789 mydomain.tld 123e4567-e89b-12d3-a456-426614174000",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsBuildDetailsV1WithResponse(ctx, args[0], args[1], uuid.MustParse(args[2]))
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
	utils.RegisterArgs(BuildCmd, utils.StringArg("username"), utils.StringArg("domain"), utils.UUIDArg("uuid"))
}
