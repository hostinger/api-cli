package actions

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <virtual-machine-id> <action-id>",
	Short: "Get action details",
	Long:  "Retrieve detailed information about a specific action performed on a specified virtual machine.\n\nUse this endpoint to monitor specific VPS operation status and details.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetActionDetailsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
