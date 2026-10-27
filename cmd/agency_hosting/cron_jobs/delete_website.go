package cron_jobs

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteCmd = &cobra.Command{
	Use:               "delete-website <website_uid> <uuid>",
	Short:             "Delete website cron job",
	Long:              "Permanently deletes the cron job identified by its uuid from an Agency Plan website.\n\nThe operation is idempotent: deleting a cron job that does not exist succeeds without error.",
	Example:           "  hostinger agency-hosting cron-jobs delete-website zpwlGlp19 01931d6f-68f5-7b72-8d9e-09c6e1e6aa0e",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete website cron job", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingDeleteWebsiteCronJobV1WithResponse(ctx, args[0], uuid.MustParse(args[1]))
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
	utils.RegisterArgs(DeleteWebsiteCmd, utils.StringArg("website_uid"), utils.UUIDArg("uuid"))
	DeleteWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
