package exposed_services

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:     "delete <instance-id> <exposed-service-id>",
	Short:   "Delete exposed service",
	Long:    "Delete a specified exposed service of a GPU instance.\n\nThe service is removed asynchronously: a successful request returns immediately and the service\nstops being reachable from the internet shortly after.\n\nUse this endpoint to close public access to a service on your GPU instance.",
	Example: "  hostinger gpu exposed-services delete 6325 3621",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().GPUGetInstancesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				a0, err := utils.ParseIntArg(args[0])
				if err != nil {
					return nil, 0, err
				}
				r, err := api.CompletionRequest().GPUGetExposedServicesV1WithResponse(ctx, a0)
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
		if err := utils.Confirm(cmd, "Delete exposed service", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUDeleteExposedServiceV1WithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(args[1]))
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
	utils.RegisterArgs(DeleteCmd, utils.IntArg("instance-id"), utils.IntArg("exposed-service-id"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
