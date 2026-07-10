package domains

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "domains",
	Short: "Domains commands",
}

func init() {
	GroupCmd.AddCommand(ChangePlanWebsiteCmd)
	GroupCmd.AddCommand(LinkToPlanWebsiteCmd)
	GroupCmd.AddCommand(ListPlanCmd)
	GroupCmd.AddCommand(UnlinkFromPlanWebsiteCmd)
}
