package ssh_keys

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListVirtualMachineCmd = &cobra.Command{
	Use:   "list-virtual-machine <virtual-machine-id>",
	Short: "List virtual machine SSH keys",
	Long:  "Retrieve SSH public keys currently configured on a specified virtual machine.\n\nOnly keys of the `root` user are listed.\n\nUse this endpoint to view SSH keys that can be used for authentication on VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSListVirtualMachineSSHKeysV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
