package catchalls

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ResendCatchAllConfirmationCmd = &cobra.Command{
	Use:   "resend-catch-all-confirmation <catchall-id>",
	Short: "Resend catch-all confirmation",
	Long:  "Resend the confirmation email to the mailbox address of an\nunconfirmed catch-all.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailResendCatchAllConfirmationV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
