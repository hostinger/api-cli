package aliases

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "aliases",
	Short: "Aliases commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateAliasCmd)
	GroupCmd.AddCommand(DeleteAliasCmd)
	GroupCmd.AddCommand(ListCmd)
}
