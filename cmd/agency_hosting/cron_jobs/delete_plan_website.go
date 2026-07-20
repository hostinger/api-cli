package cron_jobs

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeletePlanWebsiteCmd = &cobra.Command{
	Use:   "delete-plan-website <website_uid> <uuid>",
	Short: "Delete Agency Plan website cron job",
	Long:  "Permanently deletes the cron job identified by its uuid from an Agency Plan website.\n\nThe operation is idempotent: deleting a cron job that does not exist succeeds without error.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingDeleteAgencyPlanWebsiteCronJobV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
