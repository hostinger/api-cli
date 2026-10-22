package portfolio

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisablePrivacyProtectionCmd = &cobra.Command{
	Use:     "disable-privacy-protection <domain>",
	Short:   "Disable privacy protection",
	Long:    "Disable privacy protection for the domain.\n\nWhen privacy protection is disabled, domain owner's personal information is visible in public WHOIS database.\n\nUse this endpoint to make domain owner's information publicly visible.",
	Example: "  hostinger domains portfolio disable-privacy-protection mydomain.tld",
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
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Disable privacy protection", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsDisablePrivacyProtectionV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DisablePrivacyProtectionCmd, utils.StringArg("domain"))
	DisablePrivacyProtectionCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
