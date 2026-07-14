package files

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "files",
	Short: "Files commands",
}

func init() {
	GroupCmd.AddCommand(ImportPlanWebsiteFromArchiveCmd)
}
