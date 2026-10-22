package domains

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "domains",
	Short: "Domains commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteParkedCmd)
	GroupCmd.AddCommand(CreateWebsiteSubdomainCmd)
	GroupCmd.AddCommand(DeleteWebsiteParkedCmd)
	GroupCmd.AddCommand(DeleteWebsiteSubdomainCmd)
	GroupCmd.AddCommand(GenerateFreeSubdomainCmd)
	GroupCmd.AddCommand(ListWebsiteParkedCmd)
	GroupCmd.AddCommand(ListWebsiteSubdomainsCmd)
	GroupCmd.AddCommand(VerifyOwnershipCmd)
}
