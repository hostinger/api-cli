package move

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CancelOutgoingCmd = &cobra.Command{
	Use:     "cancel-outgoing <domain>",
	Short:   "Cancel outgoing domain move",
	Long:    "Cancel an outgoing move for a specified domain.\n\nThe move can only be cancelled while the receiving account has not accepted it yet.\nThe domain stays in your account.\n\nUse this endpoint to withdraw a move you no longer want to complete.",
	Example: "  hostinger domains move cancel-outgoing mydomain.tld",
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
		if err := utils.Confirm(cmd, "Cancel outgoing domain move", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsCancelOutgoingDomainMoveV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(CancelOutgoingCmd, utils.StringArg("domain"))
	CancelOutgoingCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
