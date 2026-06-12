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

var VerifyOwnershipCmd = &cobra.Command{
	Use:   "verify-ownership",
	Short: "Verify domain ownership",
	Long:  "Verify ownership of a single domain and return the verification status.\n\nUse this endpoint to check if a domain is accessible for you before using it for new websites.\nIf the domain is accessible, the response will have `is_accessible: true`.\nIf not, add the given TXT record to your domain's DNS records and try verifying again.\nKeep in mind that it may take up to 10 minutes for new TXT DNS records to propagate.\n\nSkip this verification when using Hostinger's free subdomains (*.hostingersite.com).",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(verifyOwnershipBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingVerifyDomainOwnershipV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	VerifyOwnershipCmd.Flags().StringP("domain", "", "", "Domain to verify ownership for")
	VerifyOwnershipCmd.MarkFlagRequired("domain")
}

func verifyOwnershipBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
