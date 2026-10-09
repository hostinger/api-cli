package ssh_keys

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "SSH Keys commands",
}

func init() {
	GroupCmd.AddCommand(AddVirtualMachineCmd)
	GroupCmd.AddCommand(ListVirtualMachineCmd)
	GroupCmd.AddCommand(RemoveVirtualMachineCmd)
}
