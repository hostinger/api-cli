package headless

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "headless",
	Short: "Headless commands",
}

func init() {
	GroupCmd.AddCommand(SubmitBuildReportCmd)
}
