package ssh_keys

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RemoveVirtualMachineCmd = &cobra.Command{
	Use:   "remove-virtual-machine <virtual-machine-id>",
	Short: "Remove virtual machine SSH keys",
	Long:  "Remove one or more SSH public keys from a specified virtual machine.\n\nRemoved keys can no longer be used to authenticate via SSH as the `root` user.\nReturns the remaining list of SSH keys configured on the virtual machine.\n\nUse this endpoint to revoke SSH key access to VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(removeVirtualMachineBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSRemoveVirtualMachineSSHKeysV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RemoveVirtualMachineCmd.Flags().StringSliceP("keys", "", nil, "SSH public keys in OpenSSH format to remove")
	RemoveVirtualMachineCmd.MarkFlagRequired("keys")
}

func removeVirtualMachineBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	keysVal, _ := cmd.Flags().GetStringSlice("keys")
	body["keys"] = keysVal
	return body
}
