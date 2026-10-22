package ai_content

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ai-content",
	Short: "AI Content commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(GenerateImagesCmd)
	GroupCmd.AddCommand(GenerationResultCmd)
}
