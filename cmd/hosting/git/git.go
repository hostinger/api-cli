package git

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "git",
	Short: "Git commands",
}

func init() {
	GroupCmd.AddCommand(AutoDeploymentSettingsCmd)
	GroupCmd.AddCommand(DeleteAutoDeploymentSettingsCmd)
	GroupCmd.AddCommand(DeployWebsiteRepositoryCmd)
	GroupCmd.AddCommand(GenerateSshKeyCmd)
	GroupCmd.AddCommand(ListInstallationRepositoriesCmd)
	GroupCmd.AddCommand(ListInstallationsCmd)
	GroupCmd.AddCommand(ListWebsiteRepositoriesCmd)
	GroupCmd.AddCommand(SshPublicKeyCmd)
	GroupCmd.AddCommand(UpdateAutoDeploymentSettingsCmd)
}
