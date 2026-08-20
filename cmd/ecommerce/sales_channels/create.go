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

var CreateCmd = &cobra.Command{
	Use:   "create <store_id>",
	Short: "Create a sales channel",
	Long:  "Create a sales channel for a store. A \"custom\" channel is headless: build your own frontend and keep\nyour catalog, orders, shipping and payments in sync through the Ecommerce API. A \"quick-link\" channel\nis a hosted one-page store whose handle is auto-generated.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"custom", "quick-link"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreateASalesChannelV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "", "", "Merchant-facing custom name. Required for custom channels; not supported for quick-link.")
	CreateCmd.Flags().StringP("type", "", "", "Sales channel type. \"custom\" is a headless channel: it requires a name and takes an optional public url.\n\"quick-link\" is a one-page store whose handle is auto-generated; it supports neither name nor url. (one of: custom, quick-link)")
	CreateCmd.Flags().StringP("url", "", "", "Optional public url for the channel. Custom channels only; not supported for quick-link.")
	CreateCmd.MarkFlagRequired("type")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body
}
