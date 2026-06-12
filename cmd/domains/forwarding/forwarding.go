package forwarding

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "forwarding",
	Short: "Forwarding commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
}
