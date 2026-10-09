package public_keys

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <public-key-id>",
	Short: "Delete public key",
	Long:  "Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per\nvirtual machine via `DELETE /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.\n\nDelete a public key from your account. \n\n**Deleting public key from account does not remove it from virtual machine** \n       \nUse this endpoint to remove unused SSH keys from account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeletePublicKeyV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
