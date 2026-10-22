package ai_tools

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ai-tools",
	Short: "AI Tools commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(SetOptionStatusCmd)
	GroupCmd.AddCommand(ShowOptionStatusCmd)
}
