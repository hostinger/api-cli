package recovery

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "recovery",
	Short: "Recovery commands",
}

func init() {
	GroupCmd.AddCommand(StartCmd)
	GroupCmd.AddCommand(StopCmd)
}
