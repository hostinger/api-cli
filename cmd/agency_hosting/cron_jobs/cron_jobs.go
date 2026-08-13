package cron_jobs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cron-jobs",
	Short: "Cron Jobs commands",
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
