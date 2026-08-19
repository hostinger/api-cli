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
	GroupCmd.AddCommand(BuildSettingsFromArchiveCmd)
	GroupCmd.AddCommand(CreateBuildFromArchiveCmd)
	GroupCmd.AddCommand(ListBuildsCmd)
	GroupCmd.AddCommand(ListVulnerabilitiesCmd)
	GroupCmd.AddCommand(PatchVulnerabilitiesCmd)
	GroupCmd.AddCommand(RestartApplicationCmd)
	GroupCmd.AddCommand(StartBuildCmd)
}
