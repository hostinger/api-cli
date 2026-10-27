package integrations

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ConnectAiWebsiteBuilderSiteCmd = &cobra.Command{
	Use:               "connect-ai-website-builder-site <profile-uuid>",
	Short:             "Connect AI Website Builder site",
	Long:              "Connect an AI Website Builder site the customer built manually to the profile, so its\nsubscribers flow into Reach.\n\nFails with not found when the customer owns no AI Website Builder site with that domain.\nFails with `DOMAIN_ALREADY_EXISTS` when the domain is connected to any profile - to move a\nsite to another plan, disconnect it from the old profile first - and with\n`WEBSITE_CONNECTION_ALREADY_EXISTS` when this profile already has an AI Website Builder site.\nThe eCommerce store selling on the site, if it is not connected anywhere, is connected too.\nWordPress sites and sites made with the agentic AI Website Builder cannot be connected here.",
	Example:           "  hostinger reach integrations connect-ai-website-builder-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := connectAiWebsiteBuilderSiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachConnectAIWebsiteBuilderSiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ConnectAiWebsiteBuilderSiteCmd, utils.StringArg("profile-uuid"))
	ConnectAiWebsiteBuilderSiteCmd.Flags().StringP("domain", "", "", "")
	ConnectAiWebsiteBuilderSiteCmd.MarkFlagRequired("domain")
	ConnectAiWebsiteBuilderSiteCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func connectAiWebsiteBuilderSiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
