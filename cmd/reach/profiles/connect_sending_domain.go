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

var ConnectSendingDomainCmd = &cobra.Command{
	Use:   "connect-sending-domain <profile-uuid>",
	Short: "Connect sending domain",
	Long:  "Connect a sending domain to the profile so emails can be sent from addresses on it.\n\nConnecting does not make the domain ready to send: check the DNS status endpoint next for\nrecords that are still missing. A domain the account owns or has already verified connects\nimmediately. Any other domain starts a TXT ownership verification and returns `409`;\ncomplete the verification in Reach, then retry. Returns `422` when the profile or another\nprofile already has the domain, the domain is blocked, or its ownership cannot be verified.\nAn already connected domain is never replaced.\n\nOn success the sender address of draft and scheduled campaigns, and of published automation\nand double opt-in emails, moves onto the new domain in the background, keeping the part\nbefore the `@`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(connectSendingDomainBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachConnectSendingDomainV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ConnectSendingDomainCmd.Flags().StringP("domain", "", "", "")
	ConnectSendingDomainCmd.MarkFlagRequired("domain")
}

func connectSendingDomainBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
