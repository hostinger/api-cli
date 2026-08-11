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
	Long:  "Permanently deletes an Agency Plan website. Deletion is processed asynchronously: the\nwebsite is immediately transitioned to a deleting state and the underlying server\nresources are removed in the background.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingDeleteAgencyPlanWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
