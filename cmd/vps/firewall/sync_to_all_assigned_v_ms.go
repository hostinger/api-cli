package firewall

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SyncToAllAssignedVMsCmd = &cobra.Command{
	Use:   "sync-to-all-assigned-v-ms <firewall-id>",
	Short: "Sync firewall to all assigned VMs",
	Long:  "Sync a firewall's rules to every virtual machine it's assigned to.\n\nFirewall can lose sync with a virtual machine if the firewall has new rules added, removed or updated.\n\nUse this endpoint to apply updated firewall rules to all VPS instances assigned to the firewall.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSSyncFirewallToAllAssignedVMsV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
