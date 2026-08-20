package orders

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

var FulfilCmd = &cobra.Command{
	Use:   "fulfil <store_id> <order_id>",
	Short: "Fulfil an order",
	Long:  "Create a fulfilment for the order and attach tracking in one call. Omit items to fulfil\nevery remaining unfulfilled item. Returns the updated order summary.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(fulfilBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceFulfilAnOrderV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	FulfilCmd.Flags().StringP("items", "", "", "Line items to fulfil. Omit to fulfil every remaining unfulfilled item. (JSON)")
	FulfilCmd.Flags().BoolP("notify-customer", "", false, "Whether to email the customer about the fulfilment. Defaults to true.")
	FulfilCmd.Flags().StringP("tracking-number", "", "", "Carrier tracking number for the shipment.")
	FulfilCmd.Flags().StringP("tracking-url", "", "", "Public tracking URL for the shipment. Requires tracking_number.")
}

func fulfilBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("items") {
		v, _ := cmd.Flags().GetString("items")
		body["items"] = utils.JSONValue(v, "items")
	}
	if cmd.Flags().Changed("notify-customer") {
		v, _ := cmd.Flags().GetBool("notify-customer")
		body["notify_customer"] = v
	}
	if cmd.Flags().Changed("tracking-number") {
		v, _ := cmd.Flags().GetString("tracking-number")
		body["tracking_number"] = v
	}
	if cmd.Flags().Changed("tracking-url") {
		v, _ := cmd.Flags().GetString("tracking-url")
		body["tracking_url"] = v
	}
	return body
}
