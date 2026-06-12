package records

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:   "reset <domain>",
	Short: "Reset DNS records",
	Long:  "Reset DNS zone to the default records.\n\nUse this endpoint to restore domain DNS to original configuration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(resetBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DNSResetDNSRecordsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ResetCmd.Flags().BoolP("reset-email-records", "", true, "Determines if email records should be reset")
	ResetCmd.Flags().BoolP("sync", "", true, "Determines if operation should be run synchronously")
	ResetCmd.Flags().StringSliceP("whitelisted-record-types", "", nil, "Specifies which record types to not reset")
}

func resetBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("reset-email-records") {
		v, _ := cmd.Flags().GetBool("reset-email-records")
		body["reset_email_records"] = v
	}
	if cmd.Flags().Changed("sync") {
		v, _ := cmd.Flags().GetBool("sync")
		body["sync"] = v
	}
	if cmd.Flags().Changed("whitelisted-record-types") {
		v, _ := cmd.Flags().GetStringSlice("whitelisted-record-types")
		body["whitelisted_record_types"] = v
	}
	return body
}
