package databases

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "databases",
	Short: "Databases commands",
}

func init() {
	GroupCmd.AddCommand(CreatePlanWebsiteCmd)
	GroupCmd.AddCommand(CreatePlanWebsiteUserCmd)
	GroupCmd.AddCommand(DeletePlanWebsiteCmd)
	GroupCmd.AddCommand(DeletePlanWebsiteUserCmd)
	GroupCmd.AddCommand(ListPlanWebsiteCmd)
}
