package shipping

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "shipping",
	Short: "Shipping commands",
}

func init() {
	GroupCmd.AddCommand(SetStoreCmd)
}
