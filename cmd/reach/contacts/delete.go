package contacts

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:               "delete <uuid>",
	Short:             "Delete a contact",
	Long:              "Delete a contact with the specified UUID.\n\nThis endpoint permanently removes a contact from the email marketing system.\n\nDeprecated. This endpoint cannot target a profile, so it always falls back to the\nclient's default profile and cannot delete contacts of any other profile. Use\n`DELETE /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}` instead.",
	Example:           "  hostinger reach contacts delete 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete a contact", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDeleteAContactV1WithResponse(ctx, uuid.MustParse(args[0]))
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
	utils.RegisterArgs(DeleteCmd, utils.UUIDArg("uuid"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
