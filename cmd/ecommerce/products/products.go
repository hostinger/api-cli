package products

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "products",
	Short: "Products commands",
}

func init() {
	GroupCmd.AddCommand(CreateDigitalCmd)
	GroupCmd.AddCommand(CreatePhysicalCmd)
}
