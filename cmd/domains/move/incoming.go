package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var IncomingCmd = &cobra.Command{
	Use:   "incoming <domain>",
	Short: "Get incoming domain move",
	Long:  "Retrieve the incoming move for a specified domain.\n\nReturns 404 when no account is moving this domain to you.\n\nUse this endpoint to check whether a domain addressed to you is still waiting to be accepted.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetIncomingDomainMoveV1WithResponse(context.TODO(), args[0], incomingParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	IncomingCmd.Flags().BoolP("force-sync", "", false, "Re-check the move against the registry before responding. Only has an effect while the move is in the `activating` status.")
}

func incomingParams(cmd *cobra.Command) *client.DomainsGetIncomingDomainMoveV1Params {
	params := &client.DomainsGetIncomingDomainMoveV1Params{}
	if cmd.Flags().Changed("force-sync") {
		v, _ := cmd.Flags().GetBool("force-sync")
		params.ForceSync = &v
	}
	return params
}
