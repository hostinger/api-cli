package files

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "files",
	Short: "Files commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(GenerateUploadUrlCmd)
	GroupCmd.AddCommand(ListWebsiteAndDirectoriesCmd)
	GroupCmd.AddCommand(WebsiteContentCmd)
}
