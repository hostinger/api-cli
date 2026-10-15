package checks

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "checks",
	Short: "Checks commands",
}

func init() {
	GroupCmd.AddCommand(ListCmd)
}
