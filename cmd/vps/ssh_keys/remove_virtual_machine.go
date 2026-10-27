package ssh_keys

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RemoveVirtualMachineCmd = &cobra.Command{
	Use:     "remove-virtual-machine <virtual-machine-id>",
	Short:   "Remove virtual machine SSH keys",
	Long:    "Remove one or more SSH public keys from a specified virtual machine.\n\nRemoved keys can no longer be used to authenticate via SSH as the `root` user.\nReturns the remaining list of SSH keys configured on the virtual machine.\n\nUse this endpoint to revoke SSH key access to VPS instances.",
	Example: "  hostinger vps ssh-keys remove-virtual-machine 1268054 --keys 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g... user@example.com'",
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
		body, err := removeVirtualMachineBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Remove virtual machine SSH keys", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSRemoveVirtualMachineSSHKeysV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RemoveVirtualMachineCmd, utils.IntArg("virtual-machine-id"))
	RemoveVirtualMachineCmd.Flags().StringSliceP("keys", "", nil, "SSH public keys in OpenSSH format to remove")
	RemoveVirtualMachineCmd.MarkFlagRequired("keys")
	RemoveVirtualMachineCmd.RegisterFlagCompletionFunc("keys", cobra.NoFileCompletions)
	RemoveVirtualMachineCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func removeVirtualMachineBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	keysVal, _ := cmd.Flags().GetStringSlice("keys")
	body["keys"] = keysVal
	return body, nil
}
