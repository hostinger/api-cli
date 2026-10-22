package git

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "git",
	Short: "Git commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
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
