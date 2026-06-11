package public_keys

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <public key ID>",
	Short: "Delete public key",
	Long: `This endpoint deletes a public key from your account.

Deleting public key from account does not remove it from virtual machine`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeletePublicKeyV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
