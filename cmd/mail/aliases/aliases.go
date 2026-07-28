package aliases

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "aliases",
	Short: "Aliases commands",
}

func init() {
	GroupCmd.AddCommand(CreateAliasCmd)
	GroupCmd.AddCommand(DeleteAliasCmd)
	GroupCmd.AddCommand(ListCmd)
}
