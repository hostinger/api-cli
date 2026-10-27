package php

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListVersionsForOrderCmd = &cobra.Command{
	Use:               "list-versions-for-order [order-id]",
	Short:             "List available PHP versions for an order",
	Long:              "Lists the PHP versions available to websites created under an Agency Plan order, determined by the server the order is hosted on. Use this before creating a website; for a website that already exists, call the website-scoped versions endpoint instead.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example:           "  hostinger agency-hosting php list-versions-for-order 123456",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "agency-hosting")
		if err != nil {
			return api.CommandError(err)
		}
		orderIDInt, err := utils.ScopeInt("order ID", orderID)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListAvailablePHPVersionsForAnOrderV1WithResponse(ctx, orderIDInt)
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
	utils.RegisterScopedArgs(ListVersionsForOrderCmd, 1, utils.IntArg("order_id"))
}
