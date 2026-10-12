package profiles

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteSendingDomainCmd = &cobra.Command{
	Use:   "delete-sending-domain <profile-uuid>",
	Short: "Delete sending domain",
	Long:  "Disconnect the sending domain from the profile.\n\nThe domain record is removed, so the profile has no sending domain afterwards and campaigns\ncannot be sent until a domain is connected again. Sender addresses already stored on\ncampaigns, automations and settings are left untouched and keep pointing at the removed\ndomain. A domain that is not connected to the profile also succeeds, so success is not proof\nthat anything was deleted. DNS records at the DNS provider are not removed. To switch\ndomains, delete this one and connect the new one. Confirm the domain with the customer\nbefore calling.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deleteSendingDomainBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachDeleteSendingDomainV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteSendingDomainCmd.Flags().StringP("domain", "", "", "")
	DeleteSendingDomainCmd.MarkFlagRequired("domain")
}

func deleteSendingDomainBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
