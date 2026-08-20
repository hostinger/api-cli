package discounts

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "discounts",
	Short: "Discounts commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(ListCmd)
}
