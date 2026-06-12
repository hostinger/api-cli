package virtual_machines

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ResetHostnameCmd = &cobra.Command{
	Use:   "reset-hostname <virtual-machine-id>",
	Short: "Reset hostname",
	Long:  "Reset hostname and PTR record of a specified virtual machine to default value.\n\nUse this endpoint to restore default hostname configuration for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSResetHostnameV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
