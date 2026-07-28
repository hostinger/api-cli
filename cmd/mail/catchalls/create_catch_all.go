package catchalls

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCatchAllCmd = &cobra.Command{
	Use:   "create-catch-all <mailbox-id>",
	Short: "Create catch-all",
	Long:  "Create a catch-all that routes all messages sent to unknown addresses\nof the domain to the given mailbox. The mailbox address receives a\nconfirmation email and the catch-all becomes active only after it is\nconfirmed. A domain can have only one catch-all.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailCreateCatchAllV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
