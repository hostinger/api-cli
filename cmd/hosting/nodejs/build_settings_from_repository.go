package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildSettingsFromRepositoryCmd = &cobra.Command{
	Use:     "build-settings-from-repository <username> <domain>",
	Short:   "Get Node.js build settings from repository",
	Long:    "Determine the build settings for a Node.js application from the package.json file of a branch\nin a Git repository. Use before starting a Git-source build to find out the auto-detected\nsettings instead of guessing them.\n\nThe repository is read through a Git installation: take `installation_uuid` from\n`List Git installations` and `owner` and `repository` from `List Git installation repositories`.\n\nReturns 422 when the installation is not found, the repository or branch is not found, or\nthe branch has no package.json.",
	Example: "  hostinger hosting nodejs build-settings-from-repository u123456789 mydomain.tld --owner value --repository value --branch value --installation-uuid 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
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
		if err := utils.FormatCheck(cmd, "installation-uuid", "uuid"); err != nil {
			return err
		}
		params := buildSettingsFromRepositoryParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetNodeJsBuildSettingsFromRepositoryV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(BuildSettingsFromRepositoryCmd, utils.StringArg("username"), utils.StringArg("domain"))
	BuildSettingsFromRepositoryCmd.Flags().StringP("owner", "", "", "Repository owner login. GitLab group paths use slashes.")
	BuildSettingsFromRepositoryCmd.Flags().StringP("repository", "", "", "Repository name without the .git suffix")
	BuildSettingsFromRepositoryCmd.Flags().StringP("branch", "", "", "Branch whose package.json is inspected")
	BuildSettingsFromRepositoryCmd.Flags().StringP("installation-uuid", "", "", "Git installation used to access the repository, taken from `List Git installations`")
	BuildSettingsFromRepositoryCmd.MarkFlagRequired("owner")
	BuildSettingsFromRepositoryCmd.MarkFlagRequired("repository")
	BuildSettingsFromRepositoryCmd.MarkFlagRequired("branch")
	BuildSettingsFromRepositoryCmd.MarkFlagRequired("installation-uuid")
	BuildSettingsFromRepositoryCmd.RegisterFlagCompletionFunc("owner", cobra.NoFileCompletions)
	BuildSettingsFromRepositoryCmd.RegisterFlagCompletionFunc("repository", cobra.NoFileCompletions)
	BuildSettingsFromRepositoryCmd.RegisterFlagCompletionFunc("branch", cobra.NoFileCompletions)
	BuildSettingsFromRepositoryCmd.RegisterFlagCompletionFunc("installation-uuid", cobra.NoFileCompletions)
}

func buildSettingsFromRepositoryParams(cmd *cobra.Command) *client.HostingGetNodeJsBuildSettingsFromRepositoryV1Params {
	params := &client.HostingGetNodeJsBuildSettingsFromRepositoryV1Params{}
	ownerVal, _ := cmd.Flags().GetString("owner")
	params.Owner = ownerVal
	repositoryVal, _ := cmd.Flags().GetString("repository")
	params.Repository = repositoryVal
	branchVal, _ := cmd.Flags().GetString("branch")
	params.Branch = branchVal
	installationUuidVal, _ := cmd.Flags().GetString("installation-uuid")
	params.InstallationUuid = utils.StringToUUID(installationUuidVal)
	return params
}
