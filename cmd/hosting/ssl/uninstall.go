package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:     "uninstall <username> <domain>",
	Short:   "Uninstall SSL",
	Long:    "Removes the SSL certificate assigned to the website, turns the HTTPS redirect off and cancels\na pending installation retry. The website serves plain HTTP until a new installation\ncompletes. `Get SSL status` reports `not_installed` as soon as the call returns; the call also\nsucceeds when no certificate is assigned, so repeating it is safe.\n\nReturns 422 for free subdomains (their certificate is managed by the platform) and while an\ninstallation is `installing`.",
	Example: "  hostinger hosting ssl uninstall u123456789 mydomain.tld",
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
		if err := utils.Confirm(cmd, "Uninstall SSL", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUninstallSSLV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(UninstallCmd, utils.StringArg("username"), utils.StringArg("domain"))
	UninstallCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
