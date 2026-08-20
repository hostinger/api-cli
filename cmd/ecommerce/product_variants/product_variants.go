package product_variants

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "product-variants",
	Short: "Product variants commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateInBatchCmd)
}
