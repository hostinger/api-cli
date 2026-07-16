package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RenewalInformationCmd = &cobra.Command{
	Use:   "renewal-information <domain>",
	Short: "Get domain renewal information",
	Long:  "Retrieve renewal information for a specified domain, including its status and current\nexpiration date.\n\nUse this endpoint to build renewal automation and expiry monitoring for a single domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetDomainRenewalInformationV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
