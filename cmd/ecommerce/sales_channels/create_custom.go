package sales_channels

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

var CreateCustomCmd = &cobra.Command{
	Use:   "create-custom <store_id>",
	Short: "Create custom sales channel",
	Long:  "Create a custom sales channel for a store. Build your own frontend and keep your catalog,\norders, shipping and payments in sync through the Ecommerce API.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"custom"})
		payload, err := json.Marshal(createCustomBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreateCustomSalesChannelV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCustomCmd.Flags().StringP("name", "", "", "Merchant-facing custom name shown in the sales channels list.")
	CreateCustomCmd.Flags().StringP("type", "", "", "Sales channel type. Only \"custom\" channels can be created via the API. (one of: custom)")
	CreateCustomCmd.Flags().StringP("url", "", "", "Optional public address where the custom sales channel lives.")
	CreateCustomCmd.MarkFlagRequired("name")
	CreateCustomCmd.MarkFlagRequired("type")
}

func createCustomBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body
}
