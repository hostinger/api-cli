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

var AddVirtualMachineCmd = &cobra.Command{
	Use:     "add-virtual-machine <virtual-machine-id>",
	Short:   "Add virtual machine SSH keys",
	Long:    "Add one or more SSH public keys to a specified virtual machine.\n\nKeys are added to the `root` user and can be used for passwordless SSH authentication.\nReturns the complete list of SSH keys currently configured on the virtual machine.\n\nUse this endpoint to enable SSH key authentication for VPS instances.",
	Example: "  hostinger vps ssh-keys add-virtual-machine 1268054 --keys 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g... user@example.com'",
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
		body, err := addVirtualMachineBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSAddVirtualMachineSSHKeysV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(AddVirtualMachineCmd, utils.IntArg("virtual-machine-id"))
	AddVirtualMachineCmd.Flags().StringSliceP("keys", "", nil, "SSH public keys in OpenSSH format to add")
	AddVirtualMachineCmd.MarkFlagRequired("keys")
	AddVirtualMachineCmd.RegisterFlagCompletionFunc("keys", cobra.NoFileCompletions)
}

func addVirtualMachineBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	keysVal, _ := cmd.Flags().GetStringSlice("keys")
	body["keys"] = keysVal
	return body, nil
}
