package snapshots

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var GetCmd = &cobra.Command{
	Use:   "get <domain> <snapshot ID>",
	Short: "Get DNS snapshot",
	Long:  `This endpoint retrieves a specific DNS zone snapshot for a domain, including its zone records.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DNSGetDNSSnapshotV1WithResponse(context.TODO(), args[0], utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
