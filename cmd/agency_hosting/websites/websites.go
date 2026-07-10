package websites

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "websites",
	Short: "Websites commands",
}

func init() {
	GroupCmd.AddCommand(BuildPlanNodejsAssetsCmd)
	GroupCmd.AddCommand(DeletePlanCmd)
	GroupCmd.AddCommand(ListRunningPlanProcessesCmd)
	GroupCmd.AddCommand(PlanCmd)
}
