package records

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "records",
	Short: "DNS zone record management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ResetCmd)
	GroupCmd.AddCommand(ValidateCmd)
}
