package maintenance

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "maintenance",
	Short: "Maintenance commands",
}

func init() {
	GroupCmd.AddCommand(ShowStatusCmd)
	GroupCmd.AddCommand(ToggleCmd)
}
