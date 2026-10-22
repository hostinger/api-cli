package maintenance

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "maintenance",
	Short: "Maintenance commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ShowStatusCmd)
	GroupCmd.AddCommand(ToggleCmd)
}
