package domains

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteParkedCmd = &cobra.Command{
	Use:     "delete-website-parked <username> <domain> <parked-domain>",
	Short:   "Delete website parked domain",
	Long:    "Delete an existing parked or alias domain from the selected website.\n\nUse this endpoint to remove parked domains that are no longer needed.",
	Example: "  hostinger hosting domains delete-website-parked u123456789 mydomain.tld parked-domain.com",
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
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsiteParkedDomainsV1WithResponse(ctx, args[0], args[1])
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
		if err := utils.Confirm(cmd, "Delete website parked domain", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWebsiteParkedDomainV1WithResponse(ctx, args[0], args[1], args[2])
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
	utils.RegisterArgs(DeleteWebsiteParkedCmd, utils.StringArg("username"), utils.StringArg("domain"), utils.StringArg("parked-domain"))
	DeleteWebsiteParkedCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
