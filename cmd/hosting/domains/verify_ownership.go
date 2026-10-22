package domains

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var VerifyOwnershipCmd = &cobra.Command{
	Use:               "verify-ownership",
	Short:             "Verify domain ownership",
	Long:              "Verify ownership of a single domain and return the verification status.\n\nUse this endpoint to check if a domain is accessible for you before using it for new websites.\nIf the domain is accessible, the response will have `is_accessible: true`.\nIf not, add the given TXT record to your domain's DNS records and try verifying again.\nKeep in mind that it may take up to 10 minutes for new TXT DNS records to propagate.\n\nSkip this verification when using Hostinger's free subdomains (*.hostingersite.com).",
	Example:           "  hostinger hosting domains verify-ownership --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := verifyOwnershipBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingVerifyDomainOwnershipV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(VerifyOwnershipCmd)
	VerifyOwnershipCmd.Flags().StringP("domain", "", "", "Domain to verify ownership for")
	VerifyOwnershipCmd.MarkFlagRequired("domain")
	VerifyOwnershipCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func verifyOwnershipBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
