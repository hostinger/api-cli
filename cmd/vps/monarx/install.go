package monarx

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install <virtual-machine-id>",
	Short: "Install Monarx",
	Long:  "Install the Monarx malware scanner on a specified virtual machine.\n\n[Monarx](https://www.monarx.com/) is a security tool designed to detect and\nprevent malware infections on virtual machines. By installing Monarx, users\ncan enhance the security of their virtual machines, ensuring that they are\nprotected against malicious software.\n\nUse this endpoint to enable malware protection on VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSInstallMonarxV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
