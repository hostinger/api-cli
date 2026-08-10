package move

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

var AcceptIncomingCmd = &cobra.Command{
	Use:   "accept-incoming <domain>",
	Short: "Accept incoming domain move",
	Long:  "Accept an incoming move for a specified domain.\n\nThe provided WHOIS profiles become the contacts of the domain, so they must belong\nto your account and satisfy the requirements of the TLD. Only the contact types the\ndomain actually uses are applied, but all four profile IDs have to be provided.\n\nThe move has to still be waiting for your decision, already accepted moves\ncannot be accepted again.\n\nAccepting does not complete the move. A confirmation email is sent to the email address of\nthe new owner contact, and the domain changes hands only after the change is confirmed from it.\nUntil then the move stays in the `activating` status, which can be followed with the\n[incoming move endpoint](#tag/domains-move).\n\nUse this endpoint to take ownership of a domain offered to you.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(acceptIncomingBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsAcceptIncomingDomainMoveV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AcceptIncomingCmd.Flags().StringP("domain-contacts", "", "", "WHOIS profiles of the accepting account. Only the contact types required by the TLD are applied, but all four IDs must be provided. (JSON)")
	AcceptIncomingCmd.MarkFlagRequired("domain-contacts")
}

func acceptIncomingBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainContactsVal, _ := cmd.Flags().GetString("domain-contacts")
	body["domain_contacts"] = utils.JSONValue(domainContactsVal, "domain-contacts")
	return body
}
