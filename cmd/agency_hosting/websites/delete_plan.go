package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeletePlanCmd = &cobra.Command{
	Use:   "delete-plan <website_uid>",
	Short: "Delete Agency Plan website",
	Long:  "Deletes an Agency Plan website and schedules cleanup of its resources.\n\nThis action is irreversible. Website files, databases, and linked domains are removed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingDeleteAgencyPlanWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
