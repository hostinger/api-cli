package php

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListVersionsForWebsiteCmd = &cobra.Command{
	Use:               "list-versions-for-website <website_uid>",
	Short:             "List available PHP versions for a website",
	Long:              "Lists the PHP versions an Agency Plan website can be switched to. The version the website is currently running is returned as settings.php.version by the website details endpoint.",
	Example:           "  hostinger agency-hosting php list-versions-for-website zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListAvailablePHPVersionsForAWebsiteV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListVersionsForWebsiteCmd, utils.StringArg("website_uid"))
}
