package tags

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:               "delete <profile-uuid> <tag-uuid>",
	Short:             "Delete a tag",
	Long:              "Delete a tag and remove it from every contact carrying it.\n\nThe contacts themselves are not deleted. This is idempotent: deleting a tag that does not\nexist in the profile still succeeds.",
	Example:           "  hostinger reach tags delete 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete a tag", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDeleteATagV1WithResponse(ctx, args[0], uuid.MustParse(args[1]))
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
