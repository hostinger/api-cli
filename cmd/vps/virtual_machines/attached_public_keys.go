package virtual_machines

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AttachedPublicKeysCmd = &cobra.Command{
	Use:   "attached-public-keys <virtual-machine-id>",
	Short: "Get attached public keys",
	Long:  "Retrieve public keys attached to a specified virtual machine.\n\nUse this endpoint to view SSH keys configured for specific VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetAttachedPublicKeysV1WithResponse(context.TODO(), utils.StringToInt(args[0]), attachedPublicKeysParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AttachedPublicKeysCmd.Flags().IntP("page", "", 0, "Page number")
}

func attachedPublicKeysParams(cmd *cobra.Command) *client.VPSGetAttachedPublicKeysV1Params {
	params := &client.VPSGetAttachedPublicKeysV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
