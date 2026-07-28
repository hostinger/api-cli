package forwarders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <forwarder-id>",
	Short: "Delete forwarder",
	Long:  "Delete a forwarder. The mailbox stops forwarding messages to the\ndestination address immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteForwarderV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
