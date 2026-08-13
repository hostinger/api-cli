package domains

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "domains",
	Short: "Domains commands",
}

func init() {
	GroupCmd.AddCommand(ChangeWebsiteCmd)
	GroupCmd.AddCommand(LinkToWebsiteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UnlinkFromWebsiteCmd)
}
