package cron_jobs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <username>",
	Short: "List account cron jobs",
	Long:  "Returns the list of cron jobs configured for the specified account, including their schedule and command.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListAccountCronJobsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
