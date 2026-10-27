package orders

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PlanCmd = &cobra.Command{
	Use:     "plan [order-id]",
	Short:   "Get order plan",
	Long:    "Retrieve the plan the given mail order was purchased with, including\ndomain-level and mailbox-level quotas, limits, and protocol\navailability.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail orders plan OR1a2b3c4d5e6f7g",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "id",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "mail")
		if err != nil {
			return api.CommandError(err)
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailGetOrderPlanV1WithResponse(ctx, orderID)
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
	utils.RegisterScopedArgs(PlanCmd, 1, utils.StringArg("order-id"))
}
