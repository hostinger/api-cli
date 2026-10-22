package docker

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:     "get <virtual-machine-id> <project-name>",
	Short:   "Get project contents",
	Long:    "Retrieves the complete project information including the docker-compose.yml\nfile contents, project metadata, and current deployment status.\n\nThis endpoint provides the full configuration and state details of a specific Docker Compose project.\n\nUse this to inspect project settings, review the compose file, or check the overall project health.",
	Example: "  hostinger vps docker get 1268054 my-docker-project",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetVirtualMachinesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "name",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				a0, err := utils.ParseIntArg(args[0])
				if err != nil {
					return nil, 0, err
				}
				r, err := api.CompletionRequest().VPSGetProjectListV1WithResponse(ctx, a0)
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
			r, err := api.Request().VPSGetProjectContentsV1WithResponse(ctx, utils.StringToInt(args[0]), args[1])
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
	utils.RegisterArgs(GetCmd, utils.IntArg("virtual-machine-id"), utils.StringArg("project-name"))
}
