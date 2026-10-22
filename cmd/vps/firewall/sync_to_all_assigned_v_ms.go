package firewall

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SyncToAllAssignedVMsCmd = &cobra.Command{
	Use:     "sync-to-all-assigned-v-ms <firewall-id>",
	Short:   "Sync firewall to all assigned VMs",
	Long:    "Sync a firewall's rules to every virtual machine it's assigned to.\n\nFirewall can lose sync with a virtual machine if the firewall has new rules added, removed or updated.\n\nUse this endpoint to apply updated firewall rules to all VPS instances assigned to the firewall.",
	Example: "  hostinger vps firewall sync-to-all-assigned-v-ms 9449049",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetFirewallListV1WithResponse(ctx, nil)
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
			r, err := api.Request().VPSSyncFirewallToAllAssignedVMsV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(SyncToAllAssignedVMsCmd, utils.IntArg("firewall-id"))
}
