package installations

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCoreCmd = &cobra.Command{
	Use:     "update-core <username> <software>",
	Short:   "Update WordPress core",
	Long:    "Update the WordPress core for the specified installation (minor update or a\nspecific version).\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the update\njob has been queued.",
	Example: "  hostinger wordpress installations update-core u123456789 1232456789",
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
		body, err := updateCoreBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdateWordPressCoreV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCoreCmd, utils.StringArg("username"), utils.StringArg("software"))
	UpdateCoreCmd.Flags().BoolP("minor", "", false, "Update the minor version only.")
	UpdateCoreCmd.Flags().StringP("version", "", "", "Update to a specific WordPress core version.")
	UpdateCoreCmd.RegisterFlagCompletionFunc("minor", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCoreCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
}

func updateCoreBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("minor") {
		v, _ := cmd.Flags().GetBool("minor")
		body["minor"] = v
	}
	if cmd.Flags().Changed("version") {
		v, _ := cmd.Flags().GetString("version")
		body["version"] = v
	}
	return body, nil
}
