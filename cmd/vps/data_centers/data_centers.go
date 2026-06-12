package data_centers

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "data-centers",
	Short: "Data centers commands",
}

func init() {
	GroupCmd.AddCommand(ListCmd)
}
