package databases

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "databases",
	Short: "Databases commands",
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(CreateWebsiteUserCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteUserCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
