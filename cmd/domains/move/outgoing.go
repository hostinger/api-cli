package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var OutgoingCmd = &cobra.Command{
	Use:   "outgoing <domain>",
	Short: "Get outgoing domain move",
	Long:  "Retrieve the outgoing move for a specified domain.\n\nReturns 404 when the domain has no move in progress.\n\nUse this endpoint to track the status of a move you have initiated for a single domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetOutgoingDomainMoveV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
