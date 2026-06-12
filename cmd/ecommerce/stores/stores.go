package stores

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "stores",
	Short: "Stores commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(ListCmd)
}
