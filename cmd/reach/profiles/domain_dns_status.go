package profiles

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DomainDnsStatusCmd = &cobra.Command{
	Use:               "domain-dns-status <profile-uuid>",
	Short:             "Get profile domain DNS status",
	Long:              "Retrieve the DNS configuration status for a profile's domain.\n\nThis endpoint reports the state of MX, SPF, DKIM and DMARC records, including the\nactual records found and the suggested records required for correct email delivery.",
	Example:           "  hostinger reach profiles domain-dns-status 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetProfileDomainDNSStatusV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DomainDnsStatusCmd, utils.StringArg("profile-uuid"))
}
