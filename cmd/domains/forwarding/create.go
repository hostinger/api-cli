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

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create domain forwarding",
	Long:  "Create domain forwarding configuration.\n\nUse this endpoint to set up domain redirects to other URLs.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "redirect-type", []string{"301", "302"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsCreateDomainForwardingV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("domain", "", "", "Domain name")
	CreateCmd.Flags().StringP("redirect-type", "", "", "Redirect type (one of: 301, 302)")
	CreateCmd.Flags().StringP("redirect-url", "", "", "URL to forward domain to")
	CreateCmd.MarkFlagRequired("domain")
	CreateCmd.MarkFlagRequired("redirect-type")
	CreateCmd.MarkFlagRequired("redirect-url")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	redirectTypeVal, _ := cmd.Flags().GetString("redirect-type")
	body["redirect_type"] = redirectTypeVal
	redirectUrlVal, _ := cmd.Flags().GetString("redirect-url")
	body["redirect_url"] = redirectUrlVal
	return body
}
