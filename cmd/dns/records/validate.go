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

var ValidateCmd = &cobra.Command{
	Use:   "validate <domain>",
	Short: "Validate DNS records",
	Long:  "Validate DNS records prior to update for the selected domain.\n\nIf the validation is successful, the response will contain `200 Success` code.\nIf there is validation error, the response will fail with `422 Validation error` code.\n\nUse this endpoint to verify DNS record validity before applying changes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(validateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DNSValidateDNSRecordsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ValidateCmd.Flags().BoolP("overwrite", "", true, "If `true`, resource records (RRs) matching name and type will be deleted and new RRs will be created,\notherwise resource records' ttl's are updated and new records are appended.\nIf no matching RRs are found, they are created.")
	ValidateCmd.Flags().StringP("zone", "", "", " (JSON)")
	ValidateCmd.MarkFlagRequired("zone")
}

func validateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("overwrite") {
		v, _ := cmd.Flags().GetBool("overwrite")
		body["overwrite"] = v
	}
	zoneVal, _ := cmd.Flags().GetString("zone")
	body["zone"] = utils.JSONValue(zoneVal, "zone")
	return body
}
