package payments

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableManualMethodCmd = &cobra.Command{
	Use:   "enable-manual-method <store_id>",
	Short: "Enable manual payment method",
	Long:  "Enable a manual payment method so the store can accept orders without an online payment provider.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(enableManualMethodBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceEnableManualPaymentMethodV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	EnableManualMethodCmd.Flags().StringP("title", "", "", "Optional display name shown to customers at checkout.")
}

func enableManualMethodBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body
}
