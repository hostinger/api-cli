package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisableLockCmd = &cobra.Command{
	Use:   "disable-lock <domain>",
	Short: "Disable domain lock",
	Long:  "Disable domain lock for the domain.\n\nDomain lock needs to be disabled before transferring the domain to another registrar.\n\nUse this endpoint to prepare domains for transfer to other registrars.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsDisableDomainLockV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
