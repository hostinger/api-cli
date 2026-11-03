package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PoolGroupsCmd = &cobra.Command{
	Use:               "pool-groups",
	Short:             "Get pool groups",
	Long:              "Retrieve GPU pool groups, ordered by name.\n\nEach group bundles the pools of one GPU model and lists its currently available pools,\nwhether any of them is available and the lowest hourly price among its instance types.\nThe price is returned even when the group is out of stock.\n\nUse this endpoint to show which GPU models can be ordered and from what price.",
	Example:           "  hostinger gpu catalog pool-groups",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "currency", []string{"USD", "ARS", "AMD", "AUD", "AZN", "BDT", "BRL", "BTC", "BYN", "CAD", "CHF", "CLP", "CNY", "COP", "CZK", "DKK", "EGP", "EUR", "GBP", "GEL", "HKD", "HRK", "HUF", "IDR", "ILS", "INR", "JPY", "KGS", "KRW", "KZT", "MAD", "MXN", "MYR", "NGN", "NOK", "NZD", "PEN", "PHP", "PKR", "PLN", "SEK", "SGD", "THB", "TJS", "TRY", "TWD", "UAH", "VND", "ZAR"}); err != nil {
			return err
		}
		params := poolGroupsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetPoolGroupsV1WithResponse(ctx, params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(PoolGroupsCmd)
	PoolGroupsCmd.Flags().StringP("currency", "", "", "Currency of the returned prices, `USD` by default (one of: USD, ARS, AMD, AUD, AZN, BDT, BRL, BTC, BYN, CAD, CHF, CLP, CNY, COP, CZK, DKK, EGP, EUR, GBP, GEL, HKD, HRK, HUF, IDR, ILS, INR, JPY, KGS, KRW, KZT, MAD, MXN, MYR, NGN, NOK, NZD, PEN, PHP, PKR, PLN, SEK, SGD, THB, TJS, TRY, TWD, UAH, VND, ZAR)")
	PoolGroupsCmd.RegisterFlagCompletionFunc("currency", cobra.FixedCompletions([]cobra.Completion{"USD", "ARS", "AMD", "AUD", "AZN", "BDT", "BRL", "BTC", "BYN", "CAD", "CHF", "CLP", "CNY", "COP", "CZK", "DKK", "EGP", "EUR", "GBP", "GEL", "HKD", "HRK", "HUF", "IDR", "ILS", "INR", "JPY", "KGS", "KRW", "KZT", "MAD", "MXN", "MYR", "NGN", "NOK", "NZD", "PEN", "PHP", "PKR", "PLN", "SEK", "SGD", "THB", "TJS", "TRY", "TWD", "UAH", "VND", "ZAR"}, cobra.ShellCompDirectiveNoFileComp))
}

func poolGroupsParams(cmd *cobra.Command) *client.GPUGetPoolGroupsV1Params {
	params := &client.GPUGetPoolGroupsV1Params{}
	if cmd.Flags().Changed("currency") {
		v, _ := cmd.Flags().GetString("currency")
		e := client.GPUGetPoolGroupsV1ParamsCurrency(v)
		params.Currency = &e
	}
	return params
}
