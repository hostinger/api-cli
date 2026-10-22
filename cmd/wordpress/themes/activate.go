package themes

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ActivateCmd = &cobra.Command{
	Use:     "activate <username> <software>",
	Short:   "Activate WordPress theme",
	Long:    "Activate an installed theme on a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the activation\njob has been queued.",
	Example: "  hostinger wordpress themes activate u123456789 1232456789 --theme twentytwentyone",
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
		body, err := activateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingActivateWordPressThemeV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ActivateCmd, utils.StringArg("username"), utils.StringArg("software"))
	ActivateCmd.Flags().StringP("theme", "", "", "Slug of the installed theme to activate.")
	ActivateCmd.MarkFlagRequired("theme")
	ActivateCmd.RegisterFlagCompletionFunc("theme", cobra.NoFileCompletions)
}

func activateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	themeVal, _ := cmd.Flags().GetString("theme")
	body["theme"] = themeVal
	return body, nil
}
