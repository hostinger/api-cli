package records

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <domain>",
	Short: "Delete DNS records",
	Long:  `This endpoint deletes DNS zone records for a specific domain, filtered by record name and type.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSDeleteDNSRecordsV1WithResponse(context.TODO(), args[0], zoneDeleteRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteCmd.Flags().StringArrayP("filter", "", []string{}, "Record to delete in format <name>:<type> (repeatable). Type one of: A, AAAA, ALIAS, CAA, CNAME, MX, NS, SOA, SRV, TXT")

	DeleteCmd.MarkFlagRequired("filter")
}

func zoneDeleteRequest(cmd *cobra.Command) client.DNSDeleteDNSRecordsV1JSONRequestBody {
	filterFlags, _ := cmd.Flags().GetStringArray("filter")

	body := client.DNSDeleteDNSRecordsV1JSONRequestBody{}

	for _, filter := range filterFlags {
		parts := strings.SplitN(filter, ":", 2)
		if len(parts) != 2 {
			log.Fatalf("invalid --filter %q, expected format <name>:<type>", filter)
		}

		body.Filters = append(body.Filters, struct {
			Name string                                    `json:"name"`
			Type client.DNSV1ZoneDestroyRequestFiltersType `json:"type"`
		}{
			Name: parts[0],
			Type: client.DNSV1ZoneDestroyRequestFiltersType(parts[1]),
		})
	}

	return body
}
