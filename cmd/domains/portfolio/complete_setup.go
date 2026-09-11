package portfolio

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

var CompleteSetupCmd = &cobra.Command{
	Use:   "complete-setup <domain>",
	Short: "Complete domain setup",
	Long:  "Register a domain you have already paid for but which has not been set up yet.\n\nUse this endpoint when an order completed without registering the domain, for example when\n`Purchase new domain` returned `202 Accepted` and the domain was added to your account without\nbeing registered, or when an earlier setup attempt failed. No new order is placed and no payment\nis taken: the subscription you already own is used, for the period you already paid for.\n\nA domain is left awaiting setup when the details needed to register it were missing or invalid\nas the order completed. Domains ordered elsewhere can be awaiting setup for the same reason.\nComplete the missing information, then call this endpoint. If the order itself has not completed\nyet, the domain is not on your account, wait until it appears in `Get domain list`.\n\nIf `domain_contacts` is omitted, the default WHOIS profile of that TLD is used for all four\nroles. The profile must exist and be complete for the TLD, an incomplete profile is the most\ncommon reason a domain is left awaiting setup. Create one with `Create WHOIS profile`.\n\nSome TLDs require `additional_details`. These are validated before setup, so a missing or\ninvalid value is rejected without any registration being attempted.\n\nThe domain is set up with the default nameservers and without privacy protection. Use\n`Update domain nameservers` and `Enable privacy protection` afterwards to change either.\n\nA successful response means the setup request was accepted, not that the domain is already\nregistered. Poll `Get domain list` for the outcome, the domain appears in `Get domain details`\nonly once it is registered.\n\nUse this endpoint to finish registering a domain that is awaiting setup on your account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(completeSetupBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsCompleteDomainSetupV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CompleteSetupCmd.Flags().StringP("additional-details", "", "", "Additional registration data, possible values depends on TLD (JSON)")
	CompleteSetupCmd.Flags().StringP("domain-contacts", "", "", "Domain contact information (JSON)")
}

func completeSetupBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("additional-details") {
		v, _ := cmd.Flags().GetString("additional-details")
		body["additional_details"] = utils.JSONValue(v, "additional-details")
	}
	if cmd.Flags().Changed("domain-contacts") {
		v, _ := cmd.Flags().GetString("domain-contacts")
		body["domain_contacts"] = utils.JSONValue(v, "domain-contacts")
	}
	return body
}
