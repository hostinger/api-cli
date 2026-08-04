package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CancelPendingIrtpVerificationCmd = &cobra.Command{
	Use:   "cancel-pending-irtp-verification <domain>",
	Short: "Cancel pending IRTP verification",
	Long:  "Cancel a pending IRTP verification.\n\nUse this endpoint to back out of a WHOIS change that is stuck waiting on registrant confirmation,\nfor example when the confirmation email cannot be received, without waiting out the 5-day expiry.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsCancelPendingIRTPVerificationV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
