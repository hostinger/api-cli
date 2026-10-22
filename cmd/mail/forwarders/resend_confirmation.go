package forwarders

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ResendConfirmationCmd = &cobra.Command{
	Use:               "resend-confirmation <forwarder-id>",
	Short:             "Resend forwarder confirmation",
	Long:              "Resend the confirmation email to the destination address of an\nunconfirmed forwarder.",
	Example:           "  hostinger mail forwarders resend-confirmation FW1a2b3c4d5e6f7g",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailResendForwarderConfirmationV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ResendConfirmationCmd, utils.StringArg("forwarder-id"))
}
