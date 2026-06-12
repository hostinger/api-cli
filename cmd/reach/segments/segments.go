package segments

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "segments",
	Short: "Segments commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListContactsCmd)
	GroupCmd.AddCommand(ListProfileContactsCmd)
}
