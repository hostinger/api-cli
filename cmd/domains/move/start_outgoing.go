package move

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var StartOutgoingCmd = &cobra.Command{
	Use:   "start-outgoing <domain>",
	Short: "Start outgoing domain move",
	Long:  "Initiate a move of a specified domain to another Hostinger account.\n\nThe receiving account has to already exist and accept the move before the domain changes hands.\n\nThe domain must be active. The subscription it belongs to is resolved automatically,\nand the request is rejected with a 404 status code when the domain has no domain\nsubscription of its own.\n\nDomains protected by premium protection require an additional verification step,\nsuch requests are rejected with a 428 status code.\n\nUse this endpoint to hand a domain over to another Hostinger user.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(startOutgoingBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsStartOutgoingDomainMoveV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	StartOutgoingCmd.Flags().StringP("new-customer-email", "", "", "Email address of the Hostinger account receiving the domain")
	StartOutgoingCmd.MarkFlagRequired("new-customer-email")
}

func startOutgoingBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	newCustomerEmailVal, _ := cmd.Flags().GetString("new-customer-email")
	body["new_customer_email"] = newCustomerEmailVal
	return body
}
