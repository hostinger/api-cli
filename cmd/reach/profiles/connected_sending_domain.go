package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ConnectedSendingDomainCmd = &cobra.Command{
	Use:   "connected-sending-domain <profile-uuid>",
	Short: "Get connected sending domain",
	Long:  "Get the sending domain connected to the profile, its verification status and any suspended\nsender addresses.\n\nCampaigns only go out once a domain is connected and active, so this is the cheapest way to\ncheck that precondition before building one. A profile with no domain connected returns the\nsame shape with every field set to `null`. For the individual MX, SPF, DKIM and DMARC records\nbehind the status, use the DNS status endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetConnectedSendingDomainV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
