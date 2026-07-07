package installations

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "installations",
	Short: "Installations commands",
}

func init() {
	GroupCmd.AddCommand(CheckIfAreValidCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DetectCmd)
	GroupCmd.AddCommand(InstallCmd)
	GroupCmd.AddCommand(JwtTokenCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListCoreUpdatesCmd)
	GroupCmd.AddCommand(ShowCoreVersionCmd)
	GroupCmd.AddCommand(UpdateCoreCmd)
}
