package tags

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RemoveContactFromCmd = &cobra.Command{
	Use:     "remove-contact-from <profile-uuid> <tag-uuid> <contact-uuid>",
	Short:   "Remove a contact from a tag",
	Long:    "Remove a tag from a single contact.\n\nUnlike the bulk endpoint this is applied immediately rather than queued. Neither the tag\nnor the contact is deleted.",
	Example: "  hostinger reach tags remove-contact-from 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
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
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Remove a contact from a tag", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachRemoveAContactFromATagV1WithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]))
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
	utils.RegisterArgs(RemoveContactFromCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"), utils.UUIDArg("contact-uuid"))
	RemoveContactFromCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
