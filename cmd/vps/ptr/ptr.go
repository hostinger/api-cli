package ptr

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ptr",
	Short: "PTR records commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
}
