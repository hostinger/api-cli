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

var DeleteSendingDomainCmd = &cobra.Command{
	Use:               "delete-sending-domain <profile-uuid>",
	Short:             "Delete sending domain",
	Long:              "Disconnect the sending domain from the profile.\n\nThe domain record is removed, so the profile has no sending domain afterwards and campaigns\ncannot be sent until a domain is connected again. Sender addresses already stored on\ncampaigns, automations and settings are left untouched and keep pointing at the removed\ndomain. A domain that is not connected to the profile also succeeds, so success is not proof\nthat anything was deleted. DNS records at the DNS provider are not removed. To switch\ndomains, delete this one and connect the new one. Confirm the domain with the customer\nbefore calling.",
	Example:           "  hostinger reach profiles delete-sending-domain 550e8400-e09b-41d4-a716-400055000000 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := deleteSendingDomainBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete sending domain", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDeleteSendingDomainV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DeleteSendingDomainCmd, utils.StringArg("profile-uuid"))
	DeleteSendingDomainCmd.Flags().StringP("domain", "", "", "")
	DeleteSendingDomainCmd.MarkFlagRequired("domain")
	DeleteSendingDomainCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	DeleteSendingDomainCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteSendingDomainBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
