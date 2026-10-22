package recovery

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:     "start <virtual-machine-id>",
	Short:   "Start recovery mode",
	Long:    "Initiate recovery mode for a specified virtual machine.\n\nRecovery mode is a special state that allows users to perform system rescue operations,\nsuch as repairing file systems, recovering data, or troubleshooting issues that prevent the virtual machine\nfrom booting normally.\n\nVirtual machine will boot recovery disk image and original disk image will be mounted in `/mnt` directory.\n\nUse this endpoint to enable system rescue operations on VPS instances.",
	Example: "  hostinger vps recovery start 1268054 --root-password oMeNRustosIO",
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
		body, err := startBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSStartRecoveryModeV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	StartCmd.Flags().StringP("root-password", "", "", "Temporary root password for recovery mode")
	StartCmd.MarkFlagRequired("root-password")
	StartCmd.RegisterFlagCompletionFunc("root-password", cobra.NoFileCompletions)
	utils.RegisterWaitFlags(StartCmd)
}

func startBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	rootPasswordVal, _ := cmd.Flags().GetString("root-password")
	body["root_password"] = rootPasswordVal
	return body, nil
}
