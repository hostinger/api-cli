package availability

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "availability",
	Short: "Availability commands",
}

func init() {
	GroupCmd.AddCommand(CheckCmd)
}
