package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeletePlanWebsiteUserCmd = &cobra.Command{
	Use:   "delete-plan-website-user <website_uid> <database_name> <database_user_name>",
	Short: "Delete Agency Plan website database user",
	Long:  "Permanently deletes a database user from an Agency Plan website database, revoking all access it had.\n\nThe operation is idempotent: deleting a user that does not exist succeeds without error.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingDeleteAgencyPlanWebsiteDatabaseUserV1WithResponse(context.TODO(), args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
