package monarx

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "monarx",
	Short: "Malware scanner commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ScanMetricsCmd)
	GroupCmd.AddCommand(UninstallCmd)
}
