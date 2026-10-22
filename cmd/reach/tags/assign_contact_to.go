package tags

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AssignContactToCmd = &cobra.Command{
	Use:     "assign-contact-to <profile-uuid> <tag-uuid> <contact-uuid>",
	Short:   "Assign a contact to a tag",
	Long:    "Assign a tag to a single contact.\n\nUnlike the bulk endpoint this is applied immediately rather than queued. Assigning a tag\nthe contact already carries succeeds without duplicating it.",
	Example: "  hostinger reach tags assign-contact-to 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 550e8400-e29b-41d4-a716-446655440000",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachAssignAContactToATagV1WithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]))
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
	utils.RegisterArgs(AssignContactToCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"), utils.UUIDArg("contact-uuid"))
}
