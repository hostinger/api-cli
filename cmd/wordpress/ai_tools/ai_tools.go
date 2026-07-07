package ai_tools

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ai-tools",
	Short: "AI Tools commands",
}

func init() {
	GroupCmd.AddCommand(SetOptionStatusCmd)
	GroupCmd.AddCommand(ShowOptionStatusCmd)
}
