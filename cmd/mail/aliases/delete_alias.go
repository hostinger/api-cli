package aliases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteAliasCmd = &cobra.Command{
	Use:               "delete-alias <alias-id>",
	Short:             "Delete alias",
	Long:              "Delete an alias. Messages sent to the alias address are no longer\ndelivered to the mailbox.",
	Example:           "  hostinger mail aliases delete-alias AA1a2b3c4d5e6f7g",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete alias", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailDeleteAliasV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DeleteAliasCmd, utils.StringArg("alias-id"))
	DeleteAliasCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
