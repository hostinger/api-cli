package catalog

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Catalog commands",
}

func init() {
	GroupCmd.AddCommand(ListCmd)
}
