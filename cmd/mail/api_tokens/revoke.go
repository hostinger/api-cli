package api_tokens

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RevokeCmd = &cobra.Command{
	Use:               "revoke <token-id>",
	Short:             "Revoke API token",
	Long:              "Revoke an API token. The token immediately loses access to the\nHostinger Email API (https://api.mail.hostinger.com/). This action\ncannot be undone.",
	Example:           "  hostinger mail api-tokens revoke 019683f8-1234-7abc-8def-0123456789ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Revoke API token", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailRevokeAPITokenV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(RevokeCmd, utils.StringArg("token-id"))
	RevokeCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
