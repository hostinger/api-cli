package catchalls

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCatchAllCmd = &cobra.Command{
	Use:   "delete-catch-all <catchall-id>",
	Short: "Delete catch-all",
	Long:  "Delete a catch-all. Messages sent to unknown addresses of the domain\nare no longer routed to the mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteCatchAllV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
