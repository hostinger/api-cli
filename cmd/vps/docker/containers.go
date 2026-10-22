package docker

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ContainersCmd = &cobra.Command{
	Use:     "containers <virtual-machine-id> <project-name>",
	Short:   "Get project containers",
	Long:    "Retrieves a list of all containers belonging to a specific Docker Compose project on the virtual machine.\n\nThis endpoint returns detailed information about each container including\ntheir current status, port mappings, and runtime configuration.\n\nUse this to monitor the health and state of all services within your Docker Compose project.",
	Example: "  hostinger vps docker containers 1268054 my-docker-project",
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
			r, err := api.Request().VPSGetProjectContainersV1WithResponse(ctx, utils.StringToInt(args[0]), args[1])
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
	utils.RegisterArgs(ContainersCmd, utils.IntArg("virtual-machine-id"), utils.StringArg("project-name"))
}
