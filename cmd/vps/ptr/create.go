package ptr

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

var CreateCmd = &cobra.Command{
	Use:   "create <virtual-machine-id> <ip-address-id>",
	Short: "Create PTR record",
	Long:  "Create or update a PTR (Pointer) record for a specified virtual machine.\n\nUse this endpoint to configure reverse DNS lookup for VPS IP addresses.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSCreatePTRRecordV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("domain", "", "", "Pointer record domain")
	CreateCmd.MarkFlagRequired("domain")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
