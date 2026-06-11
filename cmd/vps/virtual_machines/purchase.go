package virtual_machines

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var PurchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase and setup a new virtual machine",
	Long: `This endpoint purchases a new virtual machine using the provided catalog price item and payment method, and
sets it up with the selected operating system template in the chosen data center.

The purchased virtual machine will be set for automatic renewal.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSPurchaseNewVirtualMachineV1WithResponse(context.TODO(), purchaseRequestFromFlags(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PurchaseCmd.Flags().StringP("item_id", "", "", "Catalog price item ID")
	PurchaseCmd.Flags().IntP("payment_method_id", "", -1, "Payment method ID (default payment method is used if not provided)")
	PurchaseCmd.Flags().StringArrayP("coupon", "c", []string{}, "Discount coupon code (repeatable)")

	PurchaseCmd.MarkFlagRequired("item_id")

	addSetupFlags(PurchaseCmd)
}

func purchaseRequestFromFlags(cmd *cobra.Command) client.VPSPurchaseNewVirtualMachineV1JSONRequestBody {
	itemId, _ := cmd.Flags().GetString("item_id")
	paymentMethodId, _ := cmd.Flags().GetInt("payment_method_id")
	coupons, _ := cmd.Flags().GetStringArray("coupon")

	body := client.VPSPurchaseNewVirtualMachineV1JSONRequestBody{
		ItemId:          itemId,
		PaymentMethodId: utils.IntPtrOrNil(paymentMethodId),
		Setup:           setupRequestFromFlags(cmd),
	}

	if len(coupons) > 0 {
		c := make([]interface{}, len(coupons))
		for i, coupon := range coupons {
			c[i] = coupon
		}
		body.Coupons = &c
	}

	return body
}
