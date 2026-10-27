package files

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var WebsiteContentCmd = &cobra.Command{
	Use:     "website-content [domain]",
	Short:   "Get website file content",
	Long:    "Get a single file's content, relative to a website's document root.\n\nRead-only; refuses symlinks, oversized files, non-text file types, and files identified as\ncontaining secrets (e.g. credential files) — none of these are returned by this endpoint.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting files website-content mydomain.tld --path index.php",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
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
		params := websiteContentParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetWebsiteFileContentV1WithResponse(ctx, username, domain, params)
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
	utils.RegisterScopedArgs(WebsiteContentCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	WebsiteContentCmd.Flags().StringP("path", "", "", "File path, relative to the document root.")
	WebsiteContentCmd.Flags().IntP("from-line", "", 0, "Line offset to start reading from.")
	WebsiteContentCmd.Flags().IntP("max-lines", "", 5000, "Max number of lines to return.")
	WebsiteContentCmd.MarkFlagRequired("path")
	WebsiteContentCmd.RegisterFlagCompletionFunc("path", cobra.NoFileCompletions)
	WebsiteContentCmd.RegisterFlagCompletionFunc("from-line", cobra.NoFileCompletions)
	WebsiteContentCmd.RegisterFlagCompletionFunc("max-lines", cobra.NoFileCompletions)
}

func websiteContentParams(cmd *cobra.Command) *client.HostingGetWebsiteFileContentV1Params {
	params := &client.HostingGetWebsiteFileContentV1Params{}
	pathVal, _ := cmd.Flags().GetString("path")
	params.Path = pathVal
	if cmd.Flags().Changed("from-line") {
		v, _ := cmd.Flags().GetInt("from-line")
		params.FromLine = &v
	}
	if cmd.Flags().Changed("max-lines") {
		v, _ := cmd.Flags().GetInt("max-lines")
		params.MaxLines = &v
	}
	return params
}
