package cron_jobs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cron-jobs",
	Short: "Cron Jobs commands",
}

func init() {
	GroupCmd.AddCommand(CreatePlanWebsiteCmd)
	GroupCmd.AddCommand(DeletePlanWebsiteCmd)
	GroupCmd.AddCommand(ListPlanWebsiteCmd)
}
