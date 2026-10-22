package profiles

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SettingsCmd = &cobra.Command{
	Use:               "settings <profile-uuid>",
	Short:             "Get profile settings",
	Long:              "Get the profile's saved sender identity and brand details.\n\nRead the default sender before creating a campaign instead of asking the customer, and\nreuse the brand fields when generating email content. Any field can be `null` for a profile\nthat never filled its settings in. This does not check that the sending domain is\nverified, use the connected sending domain endpoint for that.",
	Example:           "  hostinger reach profiles settings 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetProfileSettingsV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(SettingsCmd, utils.StringArg("profile-uuid"))
}
