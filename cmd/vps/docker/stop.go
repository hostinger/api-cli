package docker

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:     "stop <virtual-machine-id> <project-name>",
	Short:   "Stop project",
	Long:    "Stops all running services in a Docker Compose project while preserving\ncontainer configurations and data volumes.\n\nThis operation gracefully shuts down containers in reverse dependency order.\n\nUse this to temporarily halt a project without removing data or configurations.",
	Example: "  hostinger vps docker stop 1268054 my-docker-project",
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
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Stop project", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSStopProjectV1WithResponse(ctx, utils.StringToInt(args[0]), args[1])
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		var waitErr error
		if wait, _ := cmd.Flags().GetBool("wait"); wait && status < 300 {
			respBody, status, waitErr = utils.Wait(cmd, respBody, utils.WaitSpec{
				IDField:    "id",
				StateField: "state",
				Done:       []string{"success"},
				Failed:     []string{"error"},
			}, func(ctx context.Context, id string) ([]byte, int, error) {
				r, err := api.Request().VPSGetActionDetailsV1WithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(id))
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			})
			if respBody == nil {
				return waitErr
			}
		}
		if err := output.Format(cmd, respBody, status); err != nil {
			return err
		}
		return waitErr
	},
}

func init() {
	utils.RegisterArgs(StopCmd, utils.IntArg("virtual-machine-id"), utils.StringArg("project-name"))
	StopCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	utils.RegisterWaitFlags(StopCmd)
}
