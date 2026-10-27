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

var DisconnectAiWebsiteBuilderSiteCmd = &cobra.Command{
	Use:               "disconnect-ai-website-builder-site <profile-uuid>",
	Short:             "Disconnect AI Website Builder site",
	Long:              "Disconnect an AI Website Builder site from the profile; the site stops sending subscribers\nto Reach.\n\nTake `domain` from the list connected websites endpoint. Works for a site that was already\ndeleted in the builder. Fails with not found when the domain is not connected to this\nprofile. Contacts and any connected eCommerce store stay. The site stays disconnected until\nit is connected again, so confirm the site and the plan with the customer before calling.",
	Example:           "  hostinger reach integrations disconnect-ai-website-builder-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := disconnectAiWebsiteBuilderSiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Disconnect AI Website Builder site", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDisconnectAIWebsiteBuilderSiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DisconnectAiWebsiteBuilderSiteCmd, utils.StringArg("profile-uuid"))
	DisconnectAiWebsiteBuilderSiteCmd.Flags().StringP("domain", "", "", "")
	DisconnectAiWebsiteBuilderSiteCmd.MarkFlagRequired("domain")
	DisconnectAiWebsiteBuilderSiteCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	DisconnectAiWebsiteBuilderSiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func disconnectAiWebsiteBuilderSiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
