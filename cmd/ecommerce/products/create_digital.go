package products

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateDigitalCmd = &cobra.Command{
	Use:   "create-digital <store_id>",
	Short: "Create digital product",
	Long:  "Create a published digital product with a single variant and an optional external download link.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createDigitalBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreateDigitalProductV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateDigitalCmd.Flags().StringP("currency", "", "", "ISO 4217 currency code. Defaults to the store's default currency when omitted.")
	CreateDigitalCmd.Flags().StringP("description", "", "", "The product description.")
	CreateDigitalCmd.Flags().StringP("download-url", "", "", "Optional external download link delivered to the customer after purchase.")
	CreateDigitalCmd.Flags().StringP("name", "", "", "The product name.")
	CreateDigitalCmd.Flags().IntP("price", "", 0, "Price in the smallest currency unit (e.g. cents). Must be positive.")
	CreateDigitalCmd.MarkFlagRequired("name")
	CreateDigitalCmd.MarkFlagRequired("price")
}

func createDigitalBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("currency") {
		v, _ := cmd.Flags().GetString("currency")
		body["currency"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("download-url") {
		v, _ := cmd.Flags().GetString("download-url")
		body["download_url"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	return body
}
