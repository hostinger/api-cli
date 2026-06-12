package public_keys

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get public keys",
	Long:  "Retrieve public keys associated with your account.\n\nUse this endpoint to view available SSH keys for VPS authentication.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetPublicKeysV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
}

func listParams(cmd *cobra.Command) *client.VPSGetPublicKeysV1Params {
	params := &client.VPSGetPublicKeysV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
