package catalog

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Catalog management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
}
