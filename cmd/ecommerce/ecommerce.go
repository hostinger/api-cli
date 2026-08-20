package ecommerce

import (
	"github.com/hostinger/api-cli/cmd/ecommerce/discounts"
	"github.com/hostinger/api-cli/cmd/ecommerce/miscellaneous"
	"github.com/hostinger/api-cli/cmd/ecommerce/orders"
	"github.com/hostinger/api-cli/cmd/ecommerce/payments"
	"github.com/hostinger/api-cli/cmd/ecommerce/product_variants"
	"github.com/hostinger/api-cli/cmd/ecommerce/products"
	"github.com/hostinger/api-cli/cmd/ecommerce/sales_channels"
	"github.com/hostinger/api-cli/cmd/ecommerce/shipping"
	"github.com/hostinger/api-cli/cmd/ecommerce/stores"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ecommerce",
	Short: "Ecommerce commands",
}

func init() {
	GroupCmd.AddCommand(discounts.GroupCmd)
	GroupCmd.AddCommand(miscellaneous.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
	GroupCmd.AddCommand(payments.GroupCmd)
	GroupCmd.AddCommand(product_variants.GroupCmd)
	GroupCmd.AddCommand(products.GroupCmd)
	GroupCmd.AddCommand(sales_channels.GroupCmd)
	GroupCmd.AddCommand(shipping.GroupCmd)
	GroupCmd.AddCommand(stores.GroupCmd)
}
