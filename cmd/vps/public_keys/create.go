package public_keys

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new public key",
	Long:  `This endpoint allows you to add a new public key to your account, which can then be attached to virtual machine instances for secure access.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSCreatePublicKeyV1WithResponse(context.TODO(), publicKeyCreateRequest(cmd, args))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "", "", "Public key name")
	CreateCmd.Flags().StringP("key", "", "", "Public key content")

	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("key")
}

func publicKeyCreateRequest(cmd *cobra.Command, args []string) client.VPSCreatePublicKeyV1JSONRequestBody {
	name, _ := cmd.Flags().GetString("name")
	key, _ := cmd.Flags().GetString("key")

	return client.VPSCreatePublicKeyV1JSONRequestBody{
		Name: name,
		Key:  key,
	}
}
