package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var OutgoingListCmd = &cobra.Command{
	Use:   "outgoing-list",
	Short: "Get outgoing domain move list",
	Long:  "Retrieve all domains you are moving to other Hostinger accounts.\n\nOnly moves which have not completed yet are returned.\n\nUse this endpoint to track moves you have initiated and the accounts they are addressed to.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetOutgoingDomainMoveListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
