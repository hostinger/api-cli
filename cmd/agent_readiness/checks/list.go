package checks

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List checks",
	Long:  "List every check an agent readiness scan runs, with its name, category and weight.\\n\nUse it to show the human-readable name of each `check_id` in a scan's results.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgentReadinessListChecksV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
