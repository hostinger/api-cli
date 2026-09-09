package websites

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "websites",
	Short: "Websites commands",
}

func init() {
	GroupCmd.AddCommand(CloneCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(EditCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(PublishCmd)
}
