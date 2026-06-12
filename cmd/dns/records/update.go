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

var UpdateCmd = &cobra.Command{
	Use:   "update <domain>",
	Short: "Update DNS records",
	Long:  "Update DNS records for the selected domain.\n\nUsing `overwrite = true` will replace existing records with the provided ones. \nOtherwise existing records will be updated and new records will be added.\n\nUse this endpoint to modify domain DNS configuration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DNSUpdateDNSRecordsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().BoolP("overwrite", "", true, "If `true`, resource records (RRs) matching name and type will be deleted and new RRs will be created,\notherwise resource records' ttl's are updated and new records are appended.\nIf no matching RRs are found, they are created.")
	UpdateCmd.Flags().StringP("zone", "", "", " (JSON)")
	UpdateCmd.MarkFlagRequired("zone")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("overwrite") {
		v, _ := cmd.Flags().GetBool("overwrite")
		body["overwrite"] = v
	}
	zoneVal, _ := cmd.Flags().GetString("zone")
	body["zone"] = utils.JSONValue(zoneVal, "zone")
	return body
}
