package payments

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "payments",
	Short: "Payments commands",
}

func init() {
	GroupCmd.AddCommand(EnableManualMethodCmd)
}
