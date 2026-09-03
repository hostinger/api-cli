package templates

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "templates",
	Short: "Templates commands",
}

func init() {
	GroupCmd.AddCommand(CreateEmailCmd)
	GroupCmd.AddCommand(ListEmailCmd)
}
