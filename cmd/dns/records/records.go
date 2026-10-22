package records

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "records",
	Short: "Zone commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ResetCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(ValidateCmd)
}
