package templates

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "templates",
	Short: "Templates commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateEmailCmd)
	GroupCmd.AddCommand(EmailCmd)
	GroupCmd.AddCommand(ListEmailCmd)
}
