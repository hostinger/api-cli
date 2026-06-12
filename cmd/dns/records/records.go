package records

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "records",
	Short: "Zone commands",
}

func init() {
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ResetCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(ValidateCmd)
}
