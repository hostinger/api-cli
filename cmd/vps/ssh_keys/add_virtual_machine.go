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

var AddVirtualMachineCmd = &cobra.Command{
	Use:   "add-virtual-machine <virtual-machine-id>",
	Short: "Add virtual machine SSH keys",
	Long:  "Add one or more SSH public keys to a specified virtual machine.\n\nKeys are added to the `root` user and can be used for passwordless SSH authentication.\nReturns the complete list of SSH keys currently configured on the virtual machine.\n\nUse this endpoint to enable SSH key authentication for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(addVirtualMachineBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSAddVirtualMachineSSHKeysV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AddVirtualMachineCmd.Flags().StringSliceP("keys", "", nil, "SSH public keys in OpenSSH format to add")
	AddVirtualMachineCmd.MarkFlagRequired("keys")
}

func addVirtualMachineBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	keysVal, _ := cmd.Flags().GetStringSlice("keys")
	body["keys"] = keysVal
	return body
}
