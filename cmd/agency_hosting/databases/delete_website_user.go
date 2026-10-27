package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteUserCmd = &cobra.Command{
	Use:               "delete-website-user <website_uid> <database_name> <database_user_name>",
	Short:             "Delete website database user",
	Long:              "Permanently deletes a database user from an Agency Plan website database, revoking all access it had.\n\nThe operation is idempotent: deleting a user that does not exist succeeds without error.",
	Example:           "  hostinger agency-hosting databases delete-website-user zpwlGlp19 my_database my_user",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete website database user", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingDeleteWebsiteDatabaseUserV1WithResponse(ctx, args[0], args[1], args[2])
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
	utils.RegisterArgs(DeleteWebsiteUserCmd, utils.StringArg("website_uid"), utils.StringArg("database_name"), utils.StringArg("database_user_name"))
	DeleteWebsiteUserCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
