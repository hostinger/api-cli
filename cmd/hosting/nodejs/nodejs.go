package nodejs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "nodejs",
	Short: "NodeJS commands",
}

func init() {
	GroupCmd.AddCommand(BuildLogsCmd)
	GroupCmd.AddCommand(CreateBuildFromArchiveCmd)
	GroupCmd.AddCommand(ListBuildsCmd)
	GroupCmd.AddCommand(RestartApplicationCmd)
}
