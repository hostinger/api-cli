package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var IncomingListCmd = &cobra.Command{
	Use:   "incoming-list",
	Short: "Get incoming domain move list",
	Long:  "Retrieve all domains other Hostinger accounts are moving to your account.\n\nMoves of every status are returned, including the ones which already completed.\n\nUse this endpoint to find domains waiting for you to accept them.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetIncomingDomainMoveListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
