package git

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteAutoDeploymentSettingsCmd = &cobra.Command{
	Use:     "delete-auto-deployment-settings <username> <domain>",
	Short:   "Delete Git auto-deployment settings",
	Long:    "Removes the Git auto-deployment settings of the website. Files already deployed stay on the\nwebsite; pushes stop deploying until settings are saved again. Succeeds also when nothing is\nconfigured.",
	Example: "  hostinger hosting git delete-auto-deployment-settings u123456789 mydomain.tld",
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
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete Git auto-deployment settings", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteGitAutoDeploymentSettingsV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(DeleteAutoDeploymentSettingsCmd, utils.StringArg("username"), utils.StringArg("domain"))
	DeleteAutoDeploymentSettingsCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
