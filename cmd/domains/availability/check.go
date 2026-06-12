package availability

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check domain availability",
	Long:  "Check availability of domain names across multiple TLDs.\n\nMultiple TLDs can be checked at once.\nIf you want alternative domains with response, provide only one TLD and set `with_alternatives` to `true`.\nTLDs should be provided without leading dot (e.g. `com`, `net`, `org`).\n\nEndpoint has rate limit of 10 requests per minute.\n\nUse this endpoint to verify domain availability before purchase.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(checkBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsCheckDomainAvailabilityV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CheckCmd.Flags().StringP("domain", "", "", "Domain name (without TLD)")
	CheckCmd.Flags().StringSliceP("tlds", "", nil, "TLDs list")
	CheckCmd.Flags().BoolP("with-alternatives", "", false, "Should response include alternatives")
	CheckCmd.MarkFlagRequired("domain")
	CheckCmd.MarkFlagRequired("tlds")
}

func checkBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	tldsVal, _ := cmd.Flags().GetStringSlice("tlds")
	body["tlds"] = tldsVal
	if cmd.Flags().Changed("with-alternatives") {
		v, _ := cmd.Flags().GetBool("with-alternatives")
		body["with_alternatives"] = v
	}
	return body
}
