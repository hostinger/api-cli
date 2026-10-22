package ssh_keys

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "SSH Keys commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(AddVirtualMachineCmd)
	GroupCmd.AddCommand(ListVirtualMachineCmd)
	GroupCmd.AddCommand(RemoveVirtualMachineCmd)
}
