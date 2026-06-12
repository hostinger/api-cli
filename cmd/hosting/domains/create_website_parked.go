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

var CreateWebsiteParkedCmd = &cobra.Command{
	Use:   "create-website-parked <username> <domain>",
	Short: "Create website parked domain",
	Long:  "Create a parked or alias domain for the selected website.\n\nProvide a domain name or IP address to park on the website so it serves the same content\nas the parent domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createWebsiteParkedBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateWebsiteParkedDomainV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateWebsiteParkedCmd.Flags().StringP("parked-domain", "", "", "Domain name or IP address to park on the selected website")
	CreateWebsiteParkedCmd.MarkFlagRequired("parked-domain")
}

func createWebsiteParkedBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	parkedDomainVal, _ := cmd.Flags().GetString("parked-domain")
	body["parked_domain"] = parkedDomainVal
	return body
}
