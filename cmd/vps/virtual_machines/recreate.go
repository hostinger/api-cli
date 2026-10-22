package virtual_machines

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RecreateCmd = &cobra.Command{
	Use:     "recreate <virtual-machine-id>",
	Short:   "Recreate virtual machine",
	Long:    "Recreate a virtual machine from scratch.\n\nThe recreation process involves reinstalling the operating system and\nresetting the virtual machine to its initial state.\nSnapshots, if there are any, will be deleted.\n\nPassword Requirements\nPassword will be checked against leaked password databases.\nRequirements for the password are:\n- At least 12 characters long\n- At least one uppercase letter\n- At least one lowercase letter\n- At least one number\n- Is not leaked publicly\n\nThis operation is irreversible and will result in the loss of all data stored on the virtual machine!\n\nUse this endpoint to completely rebuild VPS instances with fresh OS installation.",
	Example: "  hostinger vps virtual-machines recreate 1268054 --template-id 1130",
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
		body, err := recreateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Recreate virtual machine", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSRecreateVirtualMachineV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RecreateCmd, utils.IntArg("virtual-machine-id"))
	RecreateCmd.Flags().StringP("panel-password", "", "", "Panel password for the panel-based OS template. If not provided, random password will be generated. If OS does not support panel_password this field will be ignored. Password will not be shown in the response.")
	RecreateCmd.Flags().StringP("password", "", "", "Root password for the virtual machine. If not provided, random password will be generated. Password will not be shown in the response.")
	RecreateCmd.Flags().IntP("post-install-script-id", "", 0, "Post-install script to execute after virtual machine was recreated")
	RecreateCmd.Flags().IntP("template-id", "", 0, "Template ID")
	RecreateCmd.MarkFlagRequired("template-id")
	RecreateCmd.RegisterFlagCompletionFunc("panel-password", cobra.NoFileCompletions)
	RecreateCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
	RecreateCmd.RegisterFlagCompletionFunc("post-install-script-id", cobra.NoFileCompletions)
	RecreateCmd.RegisterFlagCompletionFunc("template-id", cobra.NoFileCompletions)
	RecreateCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	utils.RegisterWaitFlags(RecreateCmd)
}

func recreateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("panel-password") {
		v, _ := cmd.Flags().GetString("panel-password")
		body["panel_password"] = v
	}
	if cmd.Flags().Changed("password") {
		v, _ := cmd.Flags().GetString("password")
		body["password"] = v
	}
	if cmd.Flags().Changed("post-install-script-id") {
		v, _ := cmd.Flags().GetInt("post-install-script-id")
		body["post_install_script_id"] = v
	}
	templateIdVal, _ := cmd.Flags().GetInt("template-id")
	body["template_id"] = templateIdVal
	return body, nil
}
