package autoreplies

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "autoreplies",
	Short: "Autoreplies commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
