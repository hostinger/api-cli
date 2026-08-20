package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CancelCmd = &cobra.Command{
	Use:   "cancel <store_id> <order_id>",
	Short: "Cancel an order",
	Long:  "Cancel the order and optionally email the customer. Returns the updated order summary.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(cancelBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCancelAnOrderV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CancelCmd.Flags().BoolP("notify-customer", "", false, "Whether to email the customer about the cancellation. Defaults to true.")
}

func cancelBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("notify-customer") {
		v, _ := cmd.Flags().GetBool("notify-customer")
		body["notify_customer"] = v
	}
	return body
}
