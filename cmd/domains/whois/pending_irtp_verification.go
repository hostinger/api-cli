package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PendingIrtpVerificationCmd = &cobra.Command{
	Use:   "pending-irtp-verification <domain>",
	Short: "Get pending IRTP verification",
	Long:  "Retrieve a pending IRTP verification for a domain.\n\nBoth the old and new registrant must confirm it before the WHOIS change takes effect.\n\nUse this endpoint to check the status of a WHOIS change awaiting registrant confirmation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetPendingIRTPVerificationV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
