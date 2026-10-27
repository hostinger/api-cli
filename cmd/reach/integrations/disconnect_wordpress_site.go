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

var DisconnectWordpressSiteCmd = &cobra.Command{
	Use:               "disconnect-wordpress-site <profile-uuid>",
	Short:             "Disconnect WordPress site",
	Long:              "Disconnect a WordPress site from the profile.\n\nThe Reach plugin on the site stops sending subscribers and store data, and the site's\nproducts and coupons are removed from Reach; contacts and past orders stay. Take `domain`\nfrom the list connected websites endpoint. Fails with not found when the domain is not\nconnected to this profile. This cannot be undone through the API - the customer reconnects\nfrom the WordPress plugin - so confirm the site and the plan with the customer before calling.",
	Example:           "  hostinger reach integrations disconnect-wordpress-site 550e8400-e09b-41d4-a716-400055000000 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := disconnectWordpressSiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Disconnect WordPress site", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachDisconnectWordPressSiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DisconnectWordpressSiteCmd, utils.StringArg("profile-uuid"))
	DisconnectWordpressSiteCmd.Flags().StringP("domain", "", "", "")
	DisconnectWordpressSiteCmd.MarkFlagRequired("domain")
	DisconnectWordpressSiteCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	DisconnectWordpressSiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func disconnectWordpressSiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
