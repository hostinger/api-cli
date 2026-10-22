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

var DeployWebsiteRepositoryCmd = &cobra.Command{
	Use:     "deploy-website-repository <username> <domain>",
	Short:   "Deploy website Git repository",
	Long:    "Clones a Git repository into a directory of the website, or pulls it again. An empty or missing\ndirectory gets a clone of the branch. A directory that already holds this repository and branch is\nreset to its last commit and pulled: changes made on the server to files the repository tracks are\ndiscarded, files it does not track stay. A directory that holds other files, including another\nrepository or another branch of this one, is rejected. `composer install` runs after the clone or\npull when the repository has a `composer.json`.\n\nThe call waits for the deployment and returns its log. `is_success` false means Git or composer\nfailed and the log says why. A second call for the same directory is rejected while the first is\nstill waiting for the server. If the request times out, the deployment may still finish on the\nserver; calling again later with the same repository and branch pulls.\n\nPrivate repositories need an SSH URL and the account's Git SSH key from `Generate Git SSH key`,\nadded to the repository as a deploy key.",
	Example: "  hostinger hosting git deploy-website-repository u123456789 mydomain.tld --branch main --repository-url git@github.com:octocat/my-site.git",
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
		body, err := deployWebsiteRepositoryBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeployWebsiteGitRepositoryV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DeployWebsiteRepositoryCmd, utils.StringArg("username"), utils.StringArg("domain"))
	DeployWebsiteRepositoryCmd.Flags().StringP("branch", "", "", "Branch to clone and pull")
	DeployWebsiteRepositoryCmd.Flags().StringP("directory", "", "", "Directory under the website document root, exactly as `List website Git repositories` returns it for an existing repository. Empty, null or omitted means the document root.")
	DeployWebsiteRepositoryCmd.Flags().StringP("repository-url", "", "", "Clone URL of the repository on any Git host, SSH or HTTPS. Private repositories need an SSH URL and the account's Git SSH key added to the repository as a deploy key. An HTTP or HTTPS URL with a username or token, or any URL with a password, is rejected.")
	DeployWebsiteRepositoryCmd.MarkFlagRequired("branch")
	DeployWebsiteRepositoryCmd.MarkFlagRequired("repository-url")
	DeployWebsiteRepositoryCmd.RegisterFlagCompletionFunc("branch", cobra.NoFileCompletions)
	DeployWebsiteRepositoryCmd.RegisterFlagCompletionFunc("directory", cobra.NoFileCompletions)
	DeployWebsiteRepositoryCmd.RegisterFlagCompletionFunc("repository-url", cobra.NoFileCompletions)
}

func deployWebsiteRepositoryBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	branchVal, _ := cmd.Flags().GetString("branch")
	body["branch"] = branchVal
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	repositoryUrlVal, _ := cmd.Flags().GetString("repository-url")
	body["repository_url"] = repositoryUrlVal
	return body, nil
}
