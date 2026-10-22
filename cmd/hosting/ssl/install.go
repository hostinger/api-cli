package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:     "install <username> <domain>",
	Short:   "Install SSL",
	Long:    "Requests a lifetime SSL certificate for the website. The installation runs in the background;\n`Get SSL status` reports `active` or `failed` when it ends. An `active` lifetime certificate\ndoes not block the request: a new installation is requested, which is how a certificate is\nreinstalled.\n\nReturns 422 for free subdomains (their certificate is managed by the platform), while an\ninstallation is `installing` or `waiting_for_retry`, when the website's certificate was\nrevoked (it cannot be reissued), and when an uploaded custom certificate is installed; that\none has to be uninstalled first.",
	Example: "  hostinger hosting ssl install u123456789 mydomain.tld",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingInstallSSLV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(InstallCmd, utils.StringArg("username"), utils.StringArg("domain"))
}
