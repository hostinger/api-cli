package wordpress

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "WordPress commands",
}

func init() {
	GroupCmd.AddCommand(ChangeVersionCmd)
	GroupCmd.AddCommand(ListVersionsCmd)
	GroupCmd.AddCommand(SettingsCmd)
}
