package ai_content

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ai-content",
	Short: "AI Content commands",
}

func init() {
	GroupCmd.AddCommand(GenerateImagesCmd)
}
