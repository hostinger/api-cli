package nodejs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "nodejs",
	Short: "NodeJS commands",
}

func init() {
	GroupCmd.AddCommand(AnalyseFailedBuildCmd)
	GroupCmd.AddCommand(BuildCmd)
	GroupCmd.AddCommand(BuildLogsCmd)
	GroupCmd.AddCommand(BuildSettingsCmd)
	GroupCmd.AddCommand(BuildSettingsFromArchiveCmd)
	GroupCmd.AddCommand(ClearRuntimeLogsCmd)
	GroupCmd.AddCommand(ListBuildsCmd)
	GroupCmd.AddCommand(ListEnvironmentVariablesCmd)
	GroupCmd.AddCommand(ListVulnerabilitiesCmd)
	GroupCmd.AddCommand(PatchVulnerabilitiesCmd)
	GroupCmd.AddCommand(ReplaceEnvironmentVariablesCmd)
	GroupCmd.AddCommand(RestartApplicationCmd)
	GroupCmd.AddCommand(RuntimeLogsCmd)
	GroupCmd.AddCommand(StartBuildCmd)
	GroupCmd.AddCommand(UpdateBuildSettingsCmd)
}
