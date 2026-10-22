package whois

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PendingIrtpVerificationCmd = &cobra.Command{
	Use:     "pending-irtp-verification <domain>",
	Short:   "Get pending IRTP verification",
	Long:    "Retrieve a pending IRTP verification for a domain.\n\nBoth the old and new registrant must confirm it before the WHOIS change takes effect.\n\nUse this endpoint to check the status of a WHOIS change awaiting registrant confirmation.",
	Example: "  hostinger domains whois pending-irtp-verification mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
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
			r, err := api.Request().DomainsGetPendingIRTPVerificationV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(PendingIrtpVerificationCmd, utils.StringArg("domain"))
}
