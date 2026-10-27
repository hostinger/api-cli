package subscriptions

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisableAutoRenewalCmd = &cobra.Command{
	Use:     "disable-auto-renewal <subscription-id>",
	Short:   "Disable auto-renewal",
	Long:    "Disable auto-renewal for a subscription.\n\nUse this endpoint when disable auto-renewal for a subscription.",
	Example: "  hostinger billing subscriptions disable-auto-renewal Cxy353Uhl1xC54pG6",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().BillingGetSubscriptionListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Disable auto-renewal", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().BillingDisableAutoRenewalV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DisableAutoRenewalCmd, utils.StringArg("subscription-id"))
	DisableAutoRenewalCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
