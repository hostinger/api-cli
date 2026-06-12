package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableLockCmd = &cobra.Command{
	Use:   "enable-lock <domain>",
	Short: "Enable domain lock",
	Long:  "Enable domain lock for the domain.\n\nWhen domain lock is enabled,\nthe domain cannot be transferred to another registrar without first disabling the lock.\n\nUse this endpoint to secure domains against unauthorized transfers.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsEnableDomainLockV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
