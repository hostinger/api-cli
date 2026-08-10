package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RejectIncomingCmd = &cobra.Command{
	Use:   "reject-incoming <domain>",
	Short: "Reject incoming domain move",
	Long:  "Reject an incoming move for a specified domain.\n\nThe domain stays in the account which initiated the move.\nMoves you have already accepted cannot be rejected anymore.\n\nUse this endpoint to decline a domain you do not want to take over.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsRejectIncomingDomainMoveV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
