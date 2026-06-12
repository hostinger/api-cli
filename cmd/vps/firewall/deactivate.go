package firewall

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeactivateCmd = &cobra.Command{
	Use:   "deactivate <firewall-id> <virtual-machine-id>",
	Short: "Deactivate firewall",
	Long:  "Deactivate a firewall for a specified virtual machine.\n\nUse this endpoint to remove firewall protection from VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeactivateFirewallV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
