package discounts

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <store_id>",
	Short: "List discounts",
	Long:  "List a store's discounts. Filter by free text over code and name, or by disabled state.\nAmounts for fixed discounts are integers in the smallest currency unit; percentage\ndiscounts carry a whole-number value between 1 and 100.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "is-disabled", []string{"true", "false"})
		r, err := api.Request().EcommerceListDiscountsV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("q", "", "", "Free-text search over discount code and name.")
	ListCmd.Flags().StringP("is-disabled", "", "", "Filter by disabled state. (one of: true, false)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
}

func listParams(cmd *cobra.Command) *client.EcommerceListDiscountsV1Params {
	params := &client.EcommerceListDiscountsV1Params{}
	if cmd.Flags().Changed("q") {
		v, _ := cmd.Flags().GetString("q")
		params.Q = &v
	}
	if cmd.Flags().Changed("is-disabled") {
		v, _ := cmd.Flags().GetString("is-disabled")
		e := client.EcommerceListDiscountsV1ParamsIsDisabled(v)
		params.IsDisabled = &e
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
