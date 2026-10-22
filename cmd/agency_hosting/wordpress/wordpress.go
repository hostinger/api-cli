package wordpress

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "WordPress commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ChangeVersionCmd)
	GroupCmd.AddCommand(ListVersionsCmd)
	GroupCmd.AddCommand(SettingsCmd)
}
