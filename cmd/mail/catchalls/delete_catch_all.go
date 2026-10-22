package catchalls

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCatchAllCmd = &cobra.Command{
	Use:               "delete-catch-all <catchall-id>",
	Short:             "Delete catch-all",
	Long:              "Delete a catch-all. Messages sent to unknown addresses of the domain\nare no longer routed to the mailbox.",
	Example:           "  hostinger mail catchalls delete-catch-all CA1a2b3c4d5e6f7g",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete catch-all", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailDeleteCatchAllV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DeleteCatchAllCmd, utils.StringArg("catchall-id"))
	DeleteCatchAllCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
