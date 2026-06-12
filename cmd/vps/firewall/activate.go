package firewall

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ActivateCmd = &cobra.Command{
	Use:   "activate <firewall-id> <virtual-machine-id>",
	Short: "Activate firewall",
	Long:  "Activate a firewall for a specified virtual machine.\n\nOnly one firewall can be active for a virtual machine at a time.\n\nUse this endpoint to apply firewall rules to VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSActivateFirewallV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
