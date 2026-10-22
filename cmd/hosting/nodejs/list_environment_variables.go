package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListEnvironmentVariablesCmd = &cobra.Command{
	Use:     "list-environment-variables <username> <domain>",
	Short:   "List Node.js environment variables",
	Long:    "Lists the Node.js environment variables currently set for the website. Values are always\nmasked as `********` and cannot be read back through this API. Use this endpoint to see\nwhich keys are configured or to verify a change, not to read values.\n\nTo change variables, use the `Replace Node.js environment variables` endpoint. It replaces\nthe whole set, so never copy the masked values from this response into that request; send\nthe full desired set with real values taken from the project `.env` file or the user\nprompt instead.",
	Example: "  hostinger hosting nodejs list-environment-variables u123456789 mydomain.tld",
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
			r, err := api.Request().HostingListNodeJsEnvironmentVariablesV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(ListEnvironmentVariablesCmd, utils.StringArg("username"), utils.StringArg("domain"))
}
