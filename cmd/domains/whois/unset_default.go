package whois

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UnsetDefaultCmd = &cobra.Command{
	Use:     "unset-default <whois-id>",
	Short:   "Unset default WHOIS profile",
	Long:    "Unset WHOIS contact profile as default.\n\nThe profile itself is kept, it is only no longer pre-selected for its TLD.\n\nUse this endpoint to stop reusing contact information for new registrations.",
	Example: "  hostinger domains whois unset-default 564651",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DomainsGetWHOISProfileListV1WithResponse(ctx, nil)
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
		if err := utils.Confirm(cmd, "Unset default WHOIS profile", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsUnsetDefaultWHOISProfileV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(UnsetDefaultCmd, utils.IntArg("whois-id"))
	UnsetDefaultCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
