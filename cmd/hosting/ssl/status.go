package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:     "status <username> <domain>",
	Short:   "Get SSL status",
	Long:    "Returns the SSL state of the website: the certificate `status` and `provider`, whether the\ncertificate is a lifetime one managed by the platform, whether HTTP requests are redirected to\nHTTPS, when the certificate stops being valid and the last installation error.\n\n`installing` and `waiting_for_retry` mean an installation is in progress. `failed` means the\nlast installation gave up, or the website was not updated for 60 minutes while `installing`;\n`last_error` holds the reason when it is a known message, otherwise it is null. `expired`\nmeans the assigned certificate's validity has ended. `not_installed` means no certificate is\nassigned. Free subdomains use a platform-managed certificate: with no installation recorded\nthey report `active` with `provider` and `expires_at` null.",
	Example: "  hostinger hosting ssl status u123456789 mydomain.tld",
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
			r, err := api.Request().HostingGetSSLStatusV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(StatusCmd, utils.StringArg("username"), utils.StringArg("domain"))
}
