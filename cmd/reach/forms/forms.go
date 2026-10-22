package forms

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "forms",
	Short: "Forms commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(AddSubmitterTagsCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DuplicateCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RemoveSubmitterTagCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(UpdateTemplateCmd)
}
