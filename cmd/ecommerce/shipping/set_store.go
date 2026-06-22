package shipping

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var SetStoreCmd = &cobra.Command{
	Use:   "set-store <store_id>",
	Short: "Set store shipping",
	Long:  "Set the flat-rate shipping price for a store, creating the shipping zone if it does not exist yet.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setStoreBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceSetStoreShippingV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetStoreCmd.Flags().IntP("price", "", 0, "Flat shipping rate in the smallest currency unit (e.g. cents). Use 0 for free shipping.")
	SetStoreCmd.MarkFlagRequired("price")
}

func setStoreBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	return body
}
