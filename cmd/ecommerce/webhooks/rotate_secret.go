package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RotateSecretCmd = &cobra.Command{
	Use:               "rotate-secret <store_id> <endpoint_id>",
	Short:             "Rotate an ecommerce webhook secret",
	Long:              "Explicitly rotate the signing secret of an enabled endpoint on an active store. Save the new secret\nsecurely. Prior secrets retain their original 24-hour overlap deadlines. Repeated calls create\ndistinct secrets, up to five live secrets. Wait for an old secret to expire at the limit. After a lost\nresponse, explicitly rotate again; stored secrets cannot be retrieved. No daily renewal is required.\n\n409 codes: webhook_endpoint_disabled (endpoint is disabled or revoked), webhook_store_inactive,\nwebhook_secret_rotation_limit_reached (five live secrets). Errors return the upstream code in\n`message`, as `[Ecommerce:<code>] <reason>`.",
	Example:           "  hostinger ecommerce webhooks rotate-secret store_01J8Z5F8W9K8M4A7B3C2D1E0FG whep_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceRotateAnEcommerceWebhookSecretV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(RotateSecretCmd, utils.StringArg("store_id"), utils.StringArg("endpoint_id"))
}
