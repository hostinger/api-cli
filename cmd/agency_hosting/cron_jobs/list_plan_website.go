package cron_jobs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListPlanWebsiteCmd = &cobra.Command{
	Use:   "list-plan-website <website_uid>",
	Short: "List Agency Plan website cron jobs",
	Long:  "Returns a paginated list of cron jobs configured for an Agency Plan website.\n\nEach entry includes the schedule expression and the command executed on that schedule.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAgencyPlanWebsiteCronJobsV1WithResponse(context.TODO(), args[0], listPlanWebsiteParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListPlanWebsiteCmd.Flags().IntP("page", "", 0, "Page number")
	ListPlanWebsiteCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listPlanWebsiteParams(cmd *cobra.Command) *client.AgencyHostingListAgencyPlanWebsiteCronJobsV1Params {
	params := &client.AgencyHostingListAgencyPlanWebsiteCronJobsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
