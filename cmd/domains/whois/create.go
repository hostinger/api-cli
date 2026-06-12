package whois

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
	Short: "Create WHOIS profile",
	Long:  "Create WHOIS contact profile.\n\nUse this endpoint to add new contact information for domain registration.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "entity-type", []string{"individual", "organization"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsCreateWHOISProfileV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("country", "", "", "ISO 3166 2-letter country code")
	CreateCmd.Flags().StringP("entity-type", "", "", "Legal entity type (one of: individual, organization)")
	CreateCmd.Flags().StringP("tld", "", "", "TLD of the domain (without leading dot)")
	CreateCmd.Flags().StringP("tld-details", "", "", "TLD details (JSON)")
	CreateCmd.Flags().StringP("whois-details", "", "", "WHOIS details (JSON)")
	CreateCmd.MarkFlagRequired("country")
	CreateCmd.MarkFlagRequired("entity-type")
	CreateCmd.MarkFlagRequired("tld")
	CreateCmd.MarkFlagRequired("whois-details")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	countryVal, _ := cmd.Flags().GetString("country")
	body["country"] = countryVal
	entityTypeVal, _ := cmd.Flags().GetString("entity-type")
	body["entity_type"] = entityTypeVal
	tldVal, _ := cmd.Flags().GetString("tld")
	body["tld"] = tldVal
	if cmd.Flags().Changed("tld-details") {
		v, _ := cmd.Flags().GetString("tld-details")
		body["tld_details"] = utils.JSONValue(v, "tld-details")
	}
	whoisDetailsVal, _ := cmd.Flags().GetString("whois-details")
	body["whois_details"] = utils.JSONValue(whoisDetailsVal, "whois-details")
	return body
}
