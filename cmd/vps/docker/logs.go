package docker

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var LogsCmd = &cobra.Command{
	Use:     "logs <virtual-machine-id> <project-name>",
	Short:   "Get project logs",
	Long:    "Retrieves aggregated log entries from all services within a Docker Compose project.\n\nThis endpoint returns recent log output from each container, organized by service name with timestamps.\nThe response contains the last 300 log entries across all services.\n\nUse this for debugging, monitoring application behavior, and\ntroubleshooting issues across your entire project stack.",
	Example: "  hostinger vps docker logs 1268054 my-docker-project",
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
			r, err := api.Request().VPSGetProjectLogsV1WithResponse(ctx, utils.StringToInt(args[0]), args[1])
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
	utils.RegisterArgs(LogsCmd, utils.IntArg("virtual-machine-id"), utils.StringArg("project-name"))
}
