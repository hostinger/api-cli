package autoreplies

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "autoreplies",
	Short: "Autoreplies commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
