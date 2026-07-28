package autoreplies

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <autoreply-id>",
	Short: "Delete autoreply",
	Long:  "Delete the autoreply of a mailbox. The mailbox stops sending\nautomatic replies immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteAutoreplyV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
