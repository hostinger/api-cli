package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RegenerateSecretCmd = &cobra.Command{
	Use:               "regenerate-secret <webhook-id>",
	Short:             "Regenerate webhook secret",
	Long:              "Regenerate the secret of a webhook. The previous secret is\nimmediately invalidated. The new secret is returned only in this\nresponse and is sent as a bearer token with every delivery.",
	Example:           "  hostinger mail webhooks regenerate-secret 019683f8-1234-7abc-8def-0123456789ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailRegenerateWebhookSecretV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(RegenerateSecretCmd, utils.StringArg("webhook-id"))
}
