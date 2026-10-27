package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteCmd = &cobra.Command{
	Use:               "delete-website <website_uid> <database_name>",
	Short:             "Delete website database",
	Long:              "Permanently deletes a MySQL database and all its data from an Agency Plan website, including its users.\n\nThe operation is idempotent: deleting a database that does not exist succeeds without error.",
	Example:           "  hostinger agency-hosting databases delete-website zpwlGlp19 my_database",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete website database", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingDeleteWebsiteDatabaseV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(DeleteWebsiteCmd, utils.StringArg("website_uid"), utils.StringArg("database_name"))
	DeleteWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
