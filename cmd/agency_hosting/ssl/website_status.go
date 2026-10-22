package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var WebsiteStatusCmd = &cobra.Command{
	Use:     "website-status <website_uid> <domain>",
	Short:   "Get website SSL status",
	Long:    "Returns the SSL state of one domain of an Agency Plan website: the certificate `status`,\nwhether the certificate was uploaded by the customer, and when it stops being valid.\n\n`installing` means a certificate setup is running or retrying; the `ssl_setup` entry of\n`List website processes` shows the same progress. `active` means a valid certificate is in\nplace: uploaded by the customer, issued by the platform, or a lifetime certificate bought for\nthe domain. `failed` means the last setup gave up and no valid certificate is in place.\n`expired` means the certificate has run out. `not_installed` means the domain has no\ncertificate and no setup process. Returns 404 when the website or the domain does not exist.",
	Example: "  hostinger agency-hosting ssl website-status zpwlGlp19 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
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
			r, err := api.Request().AgencyHostingGetWebsiteSSLStatusV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(WebsiteStatusCmd, utils.StringArg("website_uid"), utils.StringArg("domain"))
}
