package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DomainDnsStatusCmd = &cobra.Command{
	Use:   "domain-dns-status <profile-uuid>",
	Short: "Get profile domain DNS status",
	Long:  "Retrieve the DNS configuration status for a profile's domain.\n\nThis endpoint reports the state of MX, SPF, DKIM and DMARC records, including the\nactual records found and the suggested records required for correct email delivery.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetProfileDomainDNSStatusV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
