package automations

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "automations",
	Short: "Automations commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListStepsCmd)
}
