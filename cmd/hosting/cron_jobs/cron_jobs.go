package cron_jobs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cron-jobs",
	Short: "Cron Jobs commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(OutputCmd)
}
