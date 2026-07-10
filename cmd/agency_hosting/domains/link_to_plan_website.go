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

var LinkToPlanWebsiteCmd = &cobra.Command{
	Use:   "link-to-plan-website <website_uid>",
	Short: "Link domain to Agency Plan website",
	Long:  "Links a domain to the specified Agency Plan website so it can serve traffic for that domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(linkToPlanWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingLinkDomainToAgencyPlanWebsiteV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	LinkToPlanWebsiteCmd.Flags().StringP("domain", "", "", "Fully qualified domain name to link to the website")
	LinkToPlanWebsiteCmd.MarkFlagRequired("domain")
}

func linkToPlanWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
