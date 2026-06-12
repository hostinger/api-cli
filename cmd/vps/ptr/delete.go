package ptr

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <virtual-machine-id> <ip-address-id>",
	Short: "Delete PTR record",
	Long:  "Delete a PTR (Pointer) record for a specified virtual machine.\n\nOnce deleted, reverse DNS lookups to the virtual machine's IP address will\nno longer return the previously configured hostname.\n\nUse this endpoint to remove reverse DNS configuration from VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeletePTRRecordV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
