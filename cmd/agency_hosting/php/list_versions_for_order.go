package php

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListVersionsForOrderCmd = &cobra.Command{
	Use:               "list-versions-for-order <order_id>",
	Short:             "List available PHP versions for an order",
	Long:              "Lists the PHP versions available to websites created under an Agency Plan order, determined by the server the order is hosted on. Use this before creating a website; for a website that already exists, call the website-scoped versions endpoint instead.",
	Example:           "  hostinger agency-hosting php list-versions-for-order 123456",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListAvailablePHPVersionsForAnOrderV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(ListVersionsForOrderCmd, utils.IntArg("order_id"))
}
