package php

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListOptionsForWebsiteCmd = &cobra.Command{
	Use:               "list-options-for-website <website_uid>",
	Short:             "List PHP options for a website",
	Long:              "Lists the php.ini directives that can be configured for an Agency Plan website, each with its default, the value currently in effect, and the values it accepts.",
	Example:           "  hostinger agency-hosting php list-options-for-website zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListPHPOptionsForAWebsiteV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListOptionsForWebsiteCmd, utils.StringArg("website_uid"))
}
