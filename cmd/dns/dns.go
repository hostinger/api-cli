package dns

import (
	"github.com/hostinger/api-cli/cmd/dns/records"
	"github.com/hostinger/api-cli/cmd/dns/snapshots"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(records.GroupCmd)
	GroupCmd.AddCommand(snapshots.GroupCmd)
}
