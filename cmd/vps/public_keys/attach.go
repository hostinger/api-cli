package public_keys

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

var AttachCmd = &cobra.Command{
	Use:   "attach <virtual-machine-id>",
	Short: "Attach public key",
	Long:  "Attach existing public keys from your account to a specified virtual machine.\n\nMultiple keys can be attached to a single virtual machine.\n\nUse this endpoint to enable SSH key authentication for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(attachBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSAttachPublicKeyV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AttachCmd.Flags().IntSliceP("ids", "", nil, "Public Key IDs to attach")
	AttachCmd.MarkFlagRequired("ids")
}

func attachBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	idsVal, _ := cmd.Flags().GetIntSlice("ids")
	body["ids"] = idsVal
	return body
}
