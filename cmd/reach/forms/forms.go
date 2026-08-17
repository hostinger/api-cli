package forms

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "forms",
	Short: "Forms commands",
}

func init() {
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
}
