package forwarders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ResendConfirmationCmd = &cobra.Command{
	Use:   "resend-confirmation <forwarder-id>",
	Short: "Resend forwarder confirmation",
	Long:  "Resend the confirmation email to the destination address of an\nunconfirmed forwarder.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailResendForwarderConfirmationV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
