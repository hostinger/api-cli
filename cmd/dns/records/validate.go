package records

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate <domain>",
	Short: "Validate DNS records",
	Long:  `This endpoint validates DNS records prior to updating them, without applying any changes.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSValidateDNSRecordsV1WithResponse(context.TODO(), args[0], zoneUpdateRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	addZoneUpdateFlags(ValidateCmd)
}
