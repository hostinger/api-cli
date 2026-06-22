package ecommerce

import (
	"github.com/hostinger/api-cli/cmd/ecommerce/payments"
	"github.com/hostinger/api-cli/cmd/ecommerce/products"
	"github.com/hostinger/api-cli/cmd/ecommerce/shipping"
	"github.com/hostinger/api-cli/cmd/ecommerce/stores"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ecommerce",
	Short: "Ecommerce commands",
}

func init() {
	GroupCmd.AddCommand(payments.GroupCmd)
	GroupCmd.AddCommand(products.GroupCmd)
	GroupCmd.AddCommand(shipping.GroupCmd)
	GroupCmd.AddCommand(stores.GroupCmd)
}
