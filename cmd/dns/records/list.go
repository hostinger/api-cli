package records

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list <domain>",
	Short: "Get DNS records",
	Long:  `This endpoint retrieves DNS zone records for a specific domain.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSGetDNSRecordsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
