package propagation

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CheckRecordCmd = &cobra.Command{
	Use:     "check-record [domain] <type>",
	Short:   "Check DNS record propagation",
	Long:    "Check whether a DNS record of a domain or host name has propagated to public resolvers around the world.\n\n`authoritative` shows what the domain's authoritative nameservers serve now. It is the best signal\nof the current record value. Its TTL tells how long resolvers may keep serving the old value.\n\n`vantage_points` shows what public resolvers at different locations return. `is_propagated` is true\nwhen the location returns the same records as the authoritative nameservers. For GeoDNS domains,\nwhose nameservers answer differently by region, true means the location returned a record.\nA records served from the Hostinger CDN count as propagated. When every responding authoritative\nnameserver returns no records, `is_propagated` is true.\n\nWorks for any domain or host name, either a domain or a subdomain. Read-only: it changes nothing.\n\nA 404 response with error code 2025 can also mean the check could not be completed, for example\nwhen the DNS checker is busy or timed out. It is safe to retry after a short delay.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger dns propagation check-record www.mydomain.tld A",
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
		domain, args, err := site.ResolveDomain(cmd, args, 1, "dns")
		if err != nil {
			return api.CommandError(err)
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSCheckDNSRecordPropagationV1WithResponse(ctx, domain, args[0])
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
	utils.RegisterScopedArgs(CheckRecordCmd, 1, utils.StringArg("domain"), utils.StringArg("type"))
}
