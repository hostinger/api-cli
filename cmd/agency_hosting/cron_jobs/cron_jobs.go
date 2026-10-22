package cron_jobs

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "cron-jobs",
	Short: "Cron Jobs commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
