package public_keys

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "public-keys",
	Short: "Public Keys commands",
}

func init() {
	GroupCmd.AddCommand(AttachCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
}
