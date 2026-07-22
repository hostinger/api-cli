package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeletePlanWebsiteCmd = &cobra.Command{
	Use:   "delete-plan-website <website_uid> <database_name>",
	Short: "Delete Agency Plan website database",
	Long:  "Permanently deletes a MySQL database and all its data from an Agency Plan website, including its users.\n\nThe operation is idempotent: deleting a database that does not exist succeeds without error.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingDeleteAgencyPlanWebsiteDatabaseV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
