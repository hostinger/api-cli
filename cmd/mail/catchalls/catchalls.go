package catchalls

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "catchalls",
	Short: "Catchalls commands",
}

func init() {
	GroupCmd.AddCommand(CreateCatchAllCmd)
	GroupCmd.AddCommand(DeleteCatchAllCmd)
	GroupCmd.AddCommand(ListCatchAllsCmd)
	GroupCmd.AddCommand(ResendCatchAllConfirmationCmd)
}
