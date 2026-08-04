package whois

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ChangeForCmd = &cobra.Command{
	Use:   "change-for",
	Short: "Change WHOIS profile for domain",
	Long:  "Change WHOIS contact profile for a domain.\n\nRepoints the given contact roles to a new WHOIS profile and submits the change to the registry.\nThe profile currently assigned to those roles is resolved automatically;\nthe request fails if the given roles are not all on the same profile today.\n\nChanging transfer sensitive fields on the owner contact starts an IRTP verification.\n\nThe change is processed asynchronously.\n\nUse this endpoint to move a registered domain onto different contact information.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changeForBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsChangeWHOISProfileForDomainV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangeForCmd.Flags().StringSliceP("change-for", "", nil, "Contact roles to repoint to the new WHOIS profile (one of: owner, admin, billing, tech)")
	ChangeForCmd.Flags().StringP("domain", "", "", "Domain name")
	ChangeForCmd.Flags().IntP("new-whois-id", "", 0, "WHOIS profile ID to assign to the domain")
	ChangeForCmd.MarkFlagRequired("change-for")
	ChangeForCmd.MarkFlagRequired("domain")
	ChangeForCmd.MarkFlagRequired("new-whois-id")
}

func changeForBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	changeForVal, _ := cmd.Flags().GetStringSlice("change-for")
	body["change_for"] = changeForVal
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	newWhoisIdVal, _ := cmd.Flags().GetInt("new-whois-id")
	body["new_whois_id"] = newWhoisIdVal
	return body
}
