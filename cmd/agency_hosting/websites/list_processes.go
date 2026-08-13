package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListProcessesCmd = &cobra.Command{
	Use:   "list-processes <website_uid>",
	Short: "List website processes",
	Long:  "Lists active and recently completed asynchronous processes for an Agency Plan website.\n\nEach process has a unique ID (for tracking), a type, and a status (running, completed, failed).\nPoll this endpoint after initiating async operations (SSL setup, backups, cloning) to track progress.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListWebsiteProcessesV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
