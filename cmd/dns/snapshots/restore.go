package snapshots

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <domain> <snapshot-id>",
	Short: "Restore DNS snapshot",
	Long:  "Restore DNS zone to the selected snapshot.\n\nUse this endpoint to revert domain DNS to a previous configuration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSRestoreDNSSnapshotV1WithResponse(context.TODO(), args[0], utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
