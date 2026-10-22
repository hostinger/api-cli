package wordpress

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SettingsCmd = &cobra.Command{
	Use:               "settings <website_uid>",
	Short:             "Get WordPress settings",
	Long:              "Returns the current WordPress settings for an Agency Plan website: installed core version,\nLiteSpeed Cache plugin status, object cache status, and maintenance mode status.",
	Example:           "  hostinger agency-hosting wordpress settings zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingGetWordPressSettingsV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(SettingsCmd, utils.StringArg("website_uid"))
}
