package snapshots

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:     "restore <domain> <snapshot-id>",
	Short:   "Restore DNS snapshot",
	Long:    "Restore DNS zone to the selected snapshot.\n\nUse this endpoint to revert domain DNS to a previous configuration.",
	Example: "  hostinger dns snapshots restore mydomain.tld 53513053",
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
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DNSGetDNSSnapshotListV1WithResponse(ctx, args[0])
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
		if err := utils.Confirm(cmd, "Restore DNS snapshot", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSRestoreDNSSnapshotV1WithResponse(ctx, args[0], utils.StringToInt(args[1]))
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
	utils.RegisterArgs(RestoreCmd, utils.StringArg("domain"), utils.IntArg("snapshot-id"))
	RestoreCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
