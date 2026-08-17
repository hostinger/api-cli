package automations

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "automations",
	Short: "Automations commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListStepsCmd)
}
