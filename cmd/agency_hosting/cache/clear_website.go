package cache

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ClearWebsiteCmd = &cobra.Command{
	Use:               "clear-website <website_uid>",
	Short:             "Clear website cache",
	Long:              "Clears cache for all domains associated with an Agency Plan website, including its preview domain.\n\nThis operation clears all cache types for the website.",
	Example:           "  hostinger agency-hosting cache clear-website zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Clear website cache", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingClearWebsiteCacheV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ClearWebsiteCmd, utils.StringArg("website_uid"))
	ClearWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
