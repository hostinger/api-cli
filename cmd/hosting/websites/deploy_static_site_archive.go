package websites

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

var DeployStaticSiteArchiveCmd = &cobra.Command{
	Use:     "deploy-static-site-archive [domain]",
	Short:   "Deploy static site archive",
	Long:    "Deploy a static application from an archive file.\n\nWARNING: this overwrites the website's existing contents and cannot be undone —\nverify this is intended before calling this endpoint.\n\nThis endpoint allows you to deploy a static application from an archive\nfile that has been uploaded to the website's directory.\n\nThis only works for static sites (pre-built HTML/CSS/JS with no build step). For\nNode.js applications, use `Create NodeJS build from archive` instead, or\n`Start Node.js build` if the archive is already uploaded. For WordPress sites,\nuse `Import WordPress website`.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting websites deploy-static-site-archive mydomain.tld --archive-path uploads/archive.zip",
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
		body, err := deployStaticSiteArchiveBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeployStaticSiteArchiveV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(DeployStaticSiteArchiveCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	DeployStaticSiteArchiveCmd.Flags().StringP("archive-path", "", "", "Relative path to the archive file from website root directory")
	DeployStaticSiteArchiveCmd.MarkFlagRequired("archive-path")
	DeployStaticSiteArchiveCmd.RegisterFlagCompletionFunc("archive-path", cobra.NoFileCompletions)
}

func deployStaticSiteArchiveBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	return body, nil
}
