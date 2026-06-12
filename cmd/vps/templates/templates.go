package templates

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "templates",
	Short: "OS Templates commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
}
