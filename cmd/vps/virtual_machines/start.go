package virtual_machines

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:     "start <virtual-machine-id>",
	Short:   "Start virtual machine",
	Long:    "Start a specified virtual machine.\n\nIf the virtual machine is already running, the request will still be processed without any effect.\n\nUse this endpoint to power on stopped VPS instances.",
	Example: "  hostinger vps vm start 1234567\n  hostinger vps vm start 1234567 --wait",
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
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSStartVirtualMachineV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(StartCmd, utils.IntArg("virtual-machine-id"))
	utils.RegisterWaitFlags(StartCmd)
}
