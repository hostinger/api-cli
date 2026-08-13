package cron_jobs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteCmd = &cobra.Command{
	Use:   "list-website <website_uid>",
	Short: "List website cron jobs",
	Long:  "Returns a paginated list of cron jobs configured for an Agency Plan website.\n\nEach entry includes the schedule expression and the command executed on that schedule.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListWebsiteCronJobsV1WithResponse(context.TODO(), args[0], listWebsiteParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListWebsiteCmd.Flags().IntP("page", "", 0, "Page number")
	ListWebsiteCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listWebsiteParams(cmd *cobra.Command) *client.AgencyHostingListWebsiteCronJobsV1Params {
	params := &client.AgencyHostingListWebsiteCronJobsV1Params{}
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
