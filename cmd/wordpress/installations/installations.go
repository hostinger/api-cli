package installations

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "installations",
	Short: "Installations commands",
}

func init() {
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(ListCmd)
}
