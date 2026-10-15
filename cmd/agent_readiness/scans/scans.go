package scans

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "scans",
	Short: "Scans commands",
}

func init() {
	GroupCmd.AddCommand(RunWebsiteCmd)
	GroupCmd.AddCommand(WebsiteCmd)
}
