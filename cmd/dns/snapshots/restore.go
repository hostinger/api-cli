package snapshots

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <domain> <snapshot ID>",
	Short: "Restore DNS snapshot",
	Long:  `This endpoint restores DNS zone records of a domain to a specific snapshot.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSRestoreDNSSnapshotV1WithResponse(context.TODO(), args[0], utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
