package automations

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListStepsCmd = &cobra.Command{
	Use:   "list-steps <profile-uuid> <automation-uuid>",
	Short: "List automation steps",
	Long:  "Get the workflow of an automation as a flat list of steps.\n\nThe steps form a tree rather than a straight line: follow `parent_uuid` to reconstruct the\nbranches, and use `step_order` to order the steps that share a parent. An automation with no\nsteps yet returns an empty list.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListAutomationStepsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
