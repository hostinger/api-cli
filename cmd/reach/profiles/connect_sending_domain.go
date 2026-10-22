package profiles

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ConnectSendingDomainCmd = &cobra.Command{
	Use:               "connect-sending-domain <profile-uuid>",
	Short:             "Connect sending domain",
	Long:              "Connect a sending domain to the profile so emails can be sent from addresses on it.\n\nConnecting does not make the domain ready to send: check the DNS status endpoint next for\nrecords that are still missing. A domain the account owns or has already verified connects\nimmediately. Any other domain starts a TXT ownership verification and returns `409`;\ncomplete the verification in Reach, then retry. Returns `422` when the profile or another\nprofile already has the domain, the domain is blocked, or its ownership cannot be verified.\nAn already connected domain is never replaced.\n\nOn success the sender address of draft and scheduled campaigns, and of published automation\nand double opt-in emails, moves onto the new domain in the background, keeping the part\nbefore the `@`.",
	Example:           "  hostinger reach profiles connect-sending-domain 550e8400-e09b-41d4-a716-400055000000 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := connectSendingDomainBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachConnectSendingDomainV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ConnectSendingDomainCmd, utils.StringArg("profile-uuid"))
	ConnectSendingDomainCmd.Flags().StringP("domain", "", "", "")
	ConnectSendingDomainCmd.MarkFlagRequired("domain")
	ConnectSendingDomainCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func connectSendingDomainBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
