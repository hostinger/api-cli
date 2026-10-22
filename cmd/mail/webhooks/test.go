package webhooks

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:               "test <webhook-id>",
	Short:             "Test webhook",
	Long:              "Send a test delivery to the webhook URL and return the result. Test\nrequests are rate limited upstream.",
	Example:           "  hostinger mail webhooks test 019683f8-1234-7abc-8def-0123456789ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailTestWebhookV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(TestCmd, utils.StringArg("webhook-id"))
}
