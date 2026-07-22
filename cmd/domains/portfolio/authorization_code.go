package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var AuthorizationCodeCmd = &cobra.Command{
	Use:   "authorization-code <domain>",
	Short: "Get domain authorization code",
	Long:  "Retrieve the authorization (EPP) code for a specified domain so it can be transferred\naway from Hostinger to another registrar.\n\nRequesting a new code invalidates any code retrieved previously.\n\nUse this endpoint to obtain the code required to transfer a domain to another registrar.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetDomainAuthorizationCodeV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
