package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UnlinkFromPlanWebsiteCmd = &cobra.Command{
	Use:   "unlink-from-plan-website <website_uid> <domain>",
	Short: "Unlink domain from Agency Plan website",
	Long:  "Unlinks a domain from the specified Agency Plan website.\n\nThe website stops serving traffic on this domain immediately.\n\nWebsite files and database are preserved, and any other linked domains remain accessible.\n\nIf this is the only domain on the website, unlinking leaves the website without an accessible domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingUnlinkDomainFromAgencyPlanWebsiteV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
