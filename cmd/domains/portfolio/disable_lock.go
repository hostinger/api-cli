package portfolio

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisableLockCmd = &cobra.Command{
	Use:     "disable-lock <domain>",
	Short:   "Disable domain lock",
	Long:    "Disable domain lock for the domain.\n\nDomain lock needs to be disabled before transferring the domain to another registrar.\n\nUse this endpoint to prepare domains for transfer to other registrars.",
	Example: "  hostinger domains portfolio disable-lock mydomain.tld",
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
		if err := utils.Confirm(cmd, "Disable domain lock", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsDisableDomainLockV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DisableLockCmd, utils.StringArg("domain"))
	DisableLockCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
