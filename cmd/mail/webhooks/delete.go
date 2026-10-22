package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:               "delete <webhook-id>",
	Short:             "Delete webhook",
	Long:              "Permanently delete a webhook. This action cannot be undone. After\ndeletion the URL no longer receives event notifications.",
	Example:           "  hostinger mail webhooks delete 019683f8-1234-7abc-8def-0123456789ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete webhook", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailDeleteWebhookV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("webhook-id"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
