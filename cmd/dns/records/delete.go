package records

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <domain>",
	Short: "Delete DNS records",
	Long:  "Delete DNS records for the selected domain.\n\nTo filter which records to delete, add the `name` of the record and `type` to the filter. \nMultiple filters can be provided with single request.\n\nIf you have multiple records with the same name and type, and you want to delete only part of them,\nrefer to the `Update zone records` endpoint.\n\nUse this endpoint to remove specific DNS records from domains.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deleteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DNSDeleteDNSRecordsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteCmd.Flags().StringP("filters", "", "", "Filter records for deletion (JSON)")
	DeleteCmd.MarkFlagRequired("filters")
}

func deleteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	filtersVal, _ := cmd.Flags().GetString("filters")
	body["filters"] = utils.JSONValue(filtersVal, "filters")
	return body
}
