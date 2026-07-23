package mailboxes

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteMailboxCmd = &cobra.Command{
	Use:   "delete-mailbox <mailbox-id>",
	Short: "Delete mailbox",
	Long:  "Delete a mailbox. The mailbox is soft-deleted and stays restorable\nfor a limited period before it is permanently removed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteMailboxV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
