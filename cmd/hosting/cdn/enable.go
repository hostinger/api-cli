package cdn

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var EnableCmd = &cobra.Command{
	Use:     "enable <username> <domain>",
	Short:   "Enable CDN",
	Long:    "Turns on the Hostinger CDN for the website domain, which speeds up page loads and reduces the\nload on the server. Works for the main domain, parked domains and subdomains of the account.\n\nEnabling changes the domain's DNS records and usually takes a few seconds. After that, DNS\npropagation can take a few more minutes before visitors are served through the CDN.\n\nRepeating the call is safe, for example after a timeout. When CDN is already enabled, it\napplies the DNS change again and refreshes the CDN SSL certificate.\n\nReturns 422 when the account is disabled, while the account is being moved to another\nserver (retry after the move finishes), or when the CDN provider rejects the change.\nReturns 404 when the domain is not a website on the account.",
	Example: "  hostinger hosting cdn enable u123456789 mydomain.tld",
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
			r, err := api.Request().HostingEnableCDNV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(EnableCmd, utils.StringArg("username"), utils.StringArg("domain"))
}
