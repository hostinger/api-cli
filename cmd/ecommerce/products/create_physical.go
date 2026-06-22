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

var CreatePhysicalCmd = &cobra.Command{
	Use:   "create-physical <store_id>",
	Short: "Create physical product",
	Long:  "Create a published physical product with a single variant priced in the store currency.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createPhysicalBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreatePhysicalProductV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreatePhysicalCmd.Flags().StringP("currency", "", "", "ISO 4217 currency code. Defaults to the store's default currency when omitted.")
	CreatePhysicalCmd.Flags().StringP("description", "", "", "The product description.")
	CreatePhysicalCmd.Flags().StringP("name", "", "", "The product name.")
	CreatePhysicalCmd.Flags().IntP("price", "", 0, "Price in the smallest currency unit (e.g. cents). Must be positive.")
	CreatePhysicalCmd.MarkFlagRequired("name")
	CreatePhysicalCmd.MarkFlagRequired("price")
}

func createPhysicalBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("currency") {
		v, _ := cmd.Flags().GetString("currency")
		body["currency"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	return body
}
