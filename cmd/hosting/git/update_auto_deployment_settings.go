package git

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateAutoDeploymentSettingsCmd = &cobra.Command{
	Use:     "update-auto-deployment-settings <username> <domain>",
	Short:   "Update Git auto-deployment settings",
	Long:    "Creates or replaces the Git auto-deployment settings of the website: repository, branch, the\ndirectory under the document root to deploy into, and `is_enabled`. Send the full set;\n`is_enabled` defaults to true and `directory` to the document root. `installation_uuid` must\nbe an installation from `List Git installations` that belongs to the same customer as the\nwebsite.\n\nFor PHP and static websites, saving with `is_enabled` true deploys the branch right away and\nevery later push to that branch deploys again. For Node.js and Website Builder websites saving\ndoes not clone anything. On a Node.js website start the first deploy with\n`Start Node.js build` using `source_type` `git`; pushes then trigger new builds with the build\nsettings stored for the website.",
	Example: "  hostinger hosting git update-auto-deployment-settings u123456789 mydomain.tld --branch main --installation-uuid 018f5e2a-1234-7890-abcd-1234567890ab --owner octocat --repository my-site",
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
		body, err := updateAutoDeploymentSettingsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdateGitAutoDeploymentSettingsV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateAutoDeploymentSettingsCmd, utils.StringArg("username"), utils.StringArg("domain"))
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("branch", "", "", "Branch to deploy")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("directory", "", "", "Subdirectory under the website document root to deploy into. Empty, null or omitted means the document root.")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("installation-uuid", "", "", "Active Git installation from `List Git installations`")
	UpdateAutoDeploymentSettingsCmd.Flags().BoolP("is-enabled", "", true, "Whether pushes to the branch deploy automatically")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("owner", "", "", "Repository owner login, as returned by `List Git installation repositories`. GitLab group paths use slashes.")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("repository", "", "", "Repository name without the .git suffix")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("branch")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("installation-uuid")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("owner")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("repository")
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("branch", cobra.NoFileCompletions)
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("directory", cobra.NoFileCompletions)
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("installation-uuid", cobra.NoFileCompletions)
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("is-enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("owner", cobra.NoFileCompletions)
	UpdateAutoDeploymentSettingsCmd.RegisterFlagCompletionFunc("repository", cobra.NoFileCompletions)
}

func updateAutoDeploymentSettingsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	branchVal, _ := cmd.Flags().GetString("branch")
	body["branch"] = branchVal
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	installationUuidVal, _ := cmd.Flags().GetString("installation-uuid")
	body["installation_uuid"] = installationUuidVal
	if cmd.Flags().Changed("is-enabled") {
		v, _ := cmd.Flags().GetBool("is-enabled")
		body["is_enabled"] = v
	}
	ownerVal, _ := cmd.Flags().GetString("owner")
	body["owner"] = ownerVal
	repositoryVal, _ := cmd.Flags().GetString("repository")
	body["repository"] = repositoryVal
	return body, nil
}
