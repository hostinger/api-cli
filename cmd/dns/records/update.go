package records

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <domain>",
	Short: "Update DNS records",
	Long: `This endpoint updates DNS records for the selected domain.

Using ` + "`overwrite = true`" + ` in the request will replace existing records matching the name and type with the provided
records. When ` + "`false`" + `, provided records are appended and existing records' TTLs are updated.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSUpdateDNSRecordsV1WithResponse(context.TODO(), args[0], zoneUpdateRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	addZoneUpdateFlags(UpdateCmd)
}

// addZoneUpdateFlags registers the flags shared by the update and validate commands,
// which use the same request body.
func addZoneUpdateFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "", "", "Name of the record (use @ for wildcard name)")
	cmd.Flags().StringP("type", "", "", "Type of the record (one of: A, AAAA, ALIAS, CAA, CNAME, MX, NS, SOA, SRV, TXT)")
	cmd.Flags().StringArrayP("content", "", []string{}, "Content of the record (repeatable to assign multiple records to the name)")
	cmd.Flags().IntP("ttl", "", 0, "TTL (Time-To-Live) of the record")
	cmd.Flags().BoolP("overwrite", "", false, "Replace records matching name and type instead of appending")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("content")
}

func zoneUpdateRequest(cmd *cobra.Command) client.DNSUpdateDNSRecordsV1JSONRequestBody {
	name, _ := cmd.Flags().GetString("name")
	recordType, _ := cmd.Flags().GetString("type")
	contents, _ := cmd.Flags().GetStringArray("content")

	records := make([]struct {
		Content string `json:"content"`
	}, 0, len(contents))
	for _, content := range contents {
		records = append(records, struct {
			Content string `json:"content"`
		}{Content: content})
	}

	entry := struct {
		Name    string `json:"name"`
		Records []struct {
			Content string `json:"content"`
		} `json:"records"`
		Ttl  *int                                  `json:"ttl,omitempty"`
		Type client.DNSV1ZoneUpdateRequestZoneType `json:"type"`
	}{
		Name:    name,
		Records: records,
		Type:    client.DNSV1ZoneUpdateRequestZoneType(recordType),
	}

	if cmd.Flags().Changed("ttl") {
		ttl, _ := cmd.Flags().GetInt("ttl")
		entry.Ttl = &ttl
	}

	body := client.DNSUpdateDNSRecordsV1JSONRequestBody{
		Zone: []struct {
			Name    string `json:"name"`
			Records []struct {
				Content string `json:"content"`
			} `json:"records"`
			Ttl  *int                                  `json:"ttl,omitempty"`
			Type client.DNSV1ZoneUpdateRequestZoneType `json:"type"`
		}{entry},
	}

	if cmd.Flags().Changed("overwrite") {
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		body.Overwrite = &overwrite
	}

	return body
}
