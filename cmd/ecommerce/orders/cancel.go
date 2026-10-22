package orders

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CancelCmd = &cobra.Command{
	Use:               "cancel <store_id> <order_id>",
	Short:             "Cancel an order",
	Long:              "Cancel the order and optionally email the customer. Returns the updated order summary.",
	Example:           "  hostinger ecommerce orders cancel store_01J8Z5F8W9K8M4A7B3C2D1E0FG order_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := cancelBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCancelAnOrderV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(CancelCmd, utils.StringArg("store_id"), utils.StringArg("order_id"))
	CancelCmd.Flags().BoolP("notify-customer", "", false, "Whether to email the customer about the cancellation. Defaults to true.")
	CancelCmd.RegisterFlagCompletionFunc("notify-customer", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func cancelBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("notify-customer") {
		v, _ := cmd.Flags().GetBool("notify-customer")
		body["notify_customer"] = v
	}
	return body, nil
}
