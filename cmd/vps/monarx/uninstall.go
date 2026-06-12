package monarx

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall <virtual-machine-id>",
	Short: "Uninstall Monarx",
	Long:  "Uninstall the Monarx malware scanner on a specified virtual machine.\n\nIf Monarx is not installed, the request will still be processed without any effect.\n\nUse this endpoint to remove malware scanner from VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSUninstallMonarxV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
