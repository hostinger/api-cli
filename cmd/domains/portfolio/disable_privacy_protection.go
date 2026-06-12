package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisablePrivacyProtectionCmd = &cobra.Command{
	Use:   "disable-privacy-protection <domain>",
	Short: "Disable privacy protection",
	Long:  "Disable privacy protection for the domain.\n\nWhen privacy protection is disabled, domain owner's personal information is visible in public WHOIS database.\n\nUse this endpoint to make domain owner's information publicly visible.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsDisablePrivacyProtectionV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
