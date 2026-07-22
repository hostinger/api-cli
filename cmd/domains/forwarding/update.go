package forwarding

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
	Short: "Update domain forwarding",
	Long:  "Update domain forwarding configuration.\n\nUse this endpoint to modify existing redirect configuration for domains.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "redirect-type", []string{"301", "302"})
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsUpdateDomainForwardingV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("redirect-type", "", "", "Redirect type (one of: 301, 302)")
	UpdateCmd.Flags().StringP("redirect-url", "", "", "URL to forward domain to")
	UpdateCmd.MarkFlagRequired("redirect-type")
	UpdateCmd.MarkFlagRequired("redirect-url")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	redirectTypeVal, _ := cmd.Flags().GetString("redirect-type")
	body["redirect_type"] = redirectTypeVal
	redirectUrlVal, _ := cmd.Flags().GetString("redirect-url")
	body["redirect_url"] = redirectUrlVal
	return body
}
