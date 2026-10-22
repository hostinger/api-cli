package plugins

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateHostingerCmd = &cobra.Command{
	Use:     "update-hostinger <username> <software>",
	Short:   "Update Hostinger WordPress plugin",
	Long:    "Update a Hostinger plugin to its latest version on a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the update job\nhas been queued.",
	Example: "  hostinger wordpress plugins update-hostinger u123456789 1232456789 --slug hostinger-affiliate-plugin",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "slug", []string{"hostinger", "hostinger-ai-assistant", "hostinger-affiliate-plugin", "hostinger-easy-onboarding", "hostinger-reach"}); err != nil {
			return err
		}
		body, err := updateHostingerBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdateHostingerWordPressPluginV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateHostingerCmd, utils.StringArg("username"), utils.StringArg("software"))
	UpdateHostingerCmd.Flags().StringP("slug", "", "", "Slug of the Hostinger plugin to update to its latest version. (one of: hostinger, hostinger-ai-assistant, hostinger-affiliate-plugin, hostinger-easy-onboarding, hostinger-reach)")
	UpdateHostingerCmd.MarkFlagRequired("slug")
	UpdateHostingerCmd.RegisterFlagCompletionFunc("slug", cobra.FixedCompletions([]cobra.Completion{"hostinger", "hostinger-ai-assistant", "hostinger-affiliate-plugin", "hostinger-easy-onboarding", "hostinger-reach"}, cobra.ShellCompDirectiveNoFileComp))
}

func updateHostingerBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	return body, nil
}
