package monarx

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "monarx",
	Short: "Malware scanner commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ScanMetricsCmd)
	GroupCmd.AddCommand(UninstallCmd)
}
