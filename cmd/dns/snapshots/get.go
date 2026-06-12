package snapshots

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <domain> <snapshot-id>",
	Short: "Get DNS snapshot",
	Long:  "Retrieve particular DNS snapshot with contents of DNS zone records.\n\nUse this endpoint to view historical DNS configurations for domains.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSGetDNSSnapshotV1WithResponse(context.TODO(), args[0], utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
