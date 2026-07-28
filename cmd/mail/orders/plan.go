package orders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PlanCmd = &cobra.Command{
	Use:   "plan <order-id>",
	Short: "Get order plan",
	Long:  "Retrieve the plan the given mail order was purchased with, including\ndomain-level and mailbox-level quotas, limits, and protocol\navailability.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailGetOrderPlanV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
