package records

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ResetCmd = &cobra.Command{
	Use:   "reset <domain>",
	Short: "Reset DNS records",
	Long:  `This endpoint resets DNS zone records to the default state for a specific domain.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSResetDNSRecordsV1WithResponse(context.TODO(), args[0], zoneResetRequest(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ResetCmd.Flags().BoolP("reset-email-records", "", false, "Determines if email records should be reset")
	ResetCmd.Flags().BoolP("sync", "", false, "Determines if operation should be run synchronously")
	ResetCmd.Flags().StringArrayP("whitelisted-record-types", "", []string{}, "Record types to not reset (repeatable)")
}

func zoneResetRequest(cmd *cobra.Command) client.DNSResetDNSRecordsV1JSONRequestBody {
	body := client.DNSResetDNSRecordsV1JSONRequestBody{}

	if cmd.Flags().Changed("reset-email-records") {
		resetEmailRecords, _ := cmd.Flags().GetBool("reset-email-records")
		body.ResetEmailRecords = &resetEmailRecords
	}

	if cmd.Flags().Changed("sync") {
		sync, _ := cmd.Flags().GetBool("sync")
		body.Sync = &sync
	}

	if cmd.Flags().Changed("whitelisted-record-types") {
		whitelistedRecordTypes, _ := cmd.Flags().GetStringArray("whitelisted-record-types")
		body.WhitelistedRecordTypes = &whitelistedRecordTypes
	}

	return body
}
