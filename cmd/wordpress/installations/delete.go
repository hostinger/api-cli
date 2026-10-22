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

var DeleteCmd = &cobra.Command{
	Use:     "delete <username> <software>",
	Short:   "Delete WordPress installation",
	Long:    "Delete the specified WordPress installation, with optional file and database\nremoval. This removes all associated components including plugins, themes,\nstaging websites and any other related data.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Example: "  hostinger wordpress installations delete u123456789 1232456789",
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
		body, err := deleteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete WordPress installation", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWordPressInstallationV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("username"), utils.StringArg("software"))
	DeleteCmd.Flags().BoolP("delete-database", "", false, "Delete the installation database.")
	DeleteCmd.Flags().BoolP("delete-files", "", false, "Delete installation files from disk.")
	DeleteCmd.RegisterFlagCompletionFunc("delete-database", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	DeleteCmd.RegisterFlagCompletionFunc("delete-files", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("delete-database") {
		v, _ := cmd.Flags().GetBool("delete-database")
		body["delete_database"] = v
	}
	if cmd.Flags().Changed("delete-files") {
		v, _ := cmd.Flags().GetBool("delete-files")
		body["delete_files"] = v
	}
	return body, nil
}
