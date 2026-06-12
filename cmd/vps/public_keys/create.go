package public_keys

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create public key",
	Long:  "Add a new public key to your account.\n\nUse this endpoint to register SSH keys for VPS authentication.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSCreatePublicKeyV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("key", "", "", "")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.MarkFlagRequired("key")
	CreateCmd.MarkFlagRequired("name")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	keyVal, _ := cmd.Flags().GetString("key")
	body["key"] = keyVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
