package actions

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "actions",
	Short: "Actions commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
}
