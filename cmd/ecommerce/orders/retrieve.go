package orders

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RetrieveCmd = &cobra.Command{
	Use:               "retrieve <store_id> <order_id>",
	Short:             "Retrieve an order",
	Long:              "Retrieve one order in full: line items (each with the id the fulfil endpoint needs),\naddresses, the totals breakdown and fulfilments with tracking. Amounts are in the\nsmallest currency unit.",
	Example:           "  hostinger ecommerce orders retrieve store_01J8Z5F8W9K8M4A7B3C2D1E0FG order_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceRetrieveAnOrderV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(RetrieveCmd, utils.StringArg("store_id"), utils.StringArg("order_id"))
}
