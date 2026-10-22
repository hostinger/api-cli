package domains

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateWebsiteSubdomainCmd = &cobra.Command{
	Use:     "create-website-subdomain <username> <domain>",
	Short:   "Create website subdomain",
	Long:    "Create a new subdomain for the selected website.\n\nProvide a subdomain prefix and, optionally, a custom directory or the\nwebsite public directory to use as the subdomain root.",
	Example: "  hostinger hosting domains create-website-subdomain u123456789 mydomain.tld --subdomain blog",
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
		body, err := createWebsiteSubdomainBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreateWebsiteSubdomainV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateWebsiteSubdomainCmd, utils.StringArg("username"), utils.StringArg("domain"))
	CreateWebsiteSubdomainCmd.Flags().StringP("directory", "", "", "Directory name for the subdomain relative to the website root")
	CreateWebsiteSubdomainCmd.Flags().BoolP("is-using-public-directory", "", false, "Use the website public directory as the subdomain root directory")
	CreateWebsiteSubdomainCmd.Flags().StringP("subdomain", "", "", "Subdomain prefix to create under the selected website")
	CreateWebsiteSubdomainCmd.MarkFlagRequired("subdomain")
	CreateWebsiteSubdomainCmd.RegisterFlagCompletionFunc("directory", cobra.NoFileCompletions)
	CreateWebsiteSubdomainCmd.RegisterFlagCompletionFunc("is-using-public-directory", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	CreateWebsiteSubdomainCmd.RegisterFlagCompletionFunc("subdomain", cobra.NoFileCompletions)
}

func createWebsiteSubdomainBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	if cmd.Flags().Changed("is-using-public-directory") {
		v, _ := cmd.Flags().GetBool("is-using-public-directory")
		body["is_using_public_directory"] = v
	}
	subdomainVal, _ := cmd.Flags().GetString("subdomain")
	body["subdomain"] = subdomainVal
	return body, nil
}
