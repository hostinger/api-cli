package domains

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ChangePlanWebsiteCmd = &cobra.Command{
	Use:   "change-plan-website <website_uid> <from_domain>",
	Short: "Change Agency Plan website domain",
	Long:  "Changes the primary domain for an Agency Plan website.\n\nProvide the current domain in the path and the new domain in the request body.\nSet domain to null to revert to the temporary domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changePlanWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingChangeAgencyPlanWebsiteDomainV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangePlanWebsiteCmd.Flags().StringP("domain", "", "", "New domain to assign to the website. Set to null to revert to the temporary domain.")
	ChangePlanWebsiteCmd.MarkFlagRequired("domain")
}

func changePlanWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
