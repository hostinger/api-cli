package catchalls

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "catchalls",
	Short: "Catchalls commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCatchAllCmd)
	GroupCmd.AddCommand(DeleteCatchAllCmd)
	GroupCmd.AddCommand(ListCatchAllsCmd)
	GroupCmd.AddCommand(ResendCatchAllConfirmationCmd)
}
