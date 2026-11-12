package pocketbase

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "PocketBase commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(BackendCmd)
	GroupCmd.AddCommand(CreateBackendCmd)
	GroupCmd.AddCommand(DeleteBackendCmd)
	GroupCmd.AddCommand(ListBackendsCmd)
	GroupCmd.AddCommand(SetSuperuserCmd)
}
