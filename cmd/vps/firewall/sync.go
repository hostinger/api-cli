package firewall

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
	Use:   "sync <firewall-id> <virtual-machine-id>",
	Short: "Sync firewall",
	Long:  "Sync a firewall for a specified virtual machine.\n\nFirewall can lose sync with virtual machine if the firewall has new rules added, removed or updated.\n\nUse this endpoint to apply updated firewall rules to VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSSyncFirewallV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
