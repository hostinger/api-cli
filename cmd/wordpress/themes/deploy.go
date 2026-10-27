package themes

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:     "deploy [domain]",
	Short:   "Deploy WordPress theme",
	Long:    "Deploy a WordPress theme from an already uploaded directory.\n\nThis endpoint allows you to deploy a WordPress theme that has been uploaded to the website's directory.\nThe theme can be optionally activated after deployment.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress themes deploy mydomain.tld --slug twentytwentyone --theme-path twentytwentyone-new",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := deployBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeployWordPressThemeV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(DeployCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	DeployCmd.Flags().BoolP("is-activated", "", false, "Whether to activate the theme after deployment")
	DeployCmd.Flags().StringP("slug", "", "", "Slug of the theme")
	DeployCmd.Flags().StringP("theme-path", "", "", "Relative path to the theme directory from wp-content/themes")
	DeployCmd.MarkFlagRequired("slug")
	DeployCmd.MarkFlagRequired("theme-path")
	DeployCmd.RegisterFlagCompletionFunc("is-activated", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	DeployCmd.RegisterFlagCompletionFunc("slug", cobra.NoFileCompletions)
	DeployCmd.RegisterFlagCompletionFunc("theme-path", cobra.NoFileCompletions)
}

func deployBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("is-activated") {
		v, _ := cmd.Flags().GetBool("is-activated")
		body["is_activated"] = v
	}
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	themePathVal, _ := cmd.Flags().GetString("theme-path")
	body["theme_path"] = themePathVal
	return body, nil
}
