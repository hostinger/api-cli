package verifications

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "verifications",
	Short: "Verifications commands",
}

func init() {
	GroupCmd.AddCommand(DirectCmd)
}
