package contacts

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteProfileCmd = &cobra.Command{
	Use:     "delete-profile <profile-uuid> <contact-uuid>",
	Short:   "Delete a profile contact",
	Long:    "Permanently delete a contact from a profile.\n\nThe contact is removed together with its custom field values and tag assignments.",
	Example: "  hostinger reach contacts delete-profile 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListContactsV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete a profile contact", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDeleteAProfileContactV1WithResponse(ctx, args[0], uuid.MustParse(args[1]))
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
	utils.RegisterArgs(DeleteProfileCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("contact-uuid"))
	DeleteProfileCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
