package cron_jobs

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var OutputCmd = &cobra.Command{
	Use:   "output <username> <uid>",
	Short: "Get cron job output",
	Long:  "Returns the output captured from the last execution of the cron job identified by its uid.\n\nThe uid is returned by the list cron jobs endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetCronJobOutputV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
