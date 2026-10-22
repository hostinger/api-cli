package plugins

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CheckIfWooCommerceIsInstalledCmd = &cobra.Command{
	Use:               "check-if-woo-commerce-is-installed",
	Short:             "Check if WooCommerce is installed",
	Long:              "Check whether WooCommerce is installed on any WordPress installation of a\ndomain. Optionally filter by domain to scope the check.",
	Example:           "  hostinger wordpress plugins check-if-woo-commerce-is-installed",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := checkIfWooCommerceIsInstalledParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCheckIfWooCommerceIsInstalledV1WithResponse(ctx, params)
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
	utils.RegisterArgs(CheckIfWooCommerceIsInstalledCmd)
	CheckIfWooCommerceIsInstalledCmd.Flags().StringP("domain", "", "", "Filter by domain name (case-insensitive substring match)")
	CheckIfWooCommerceIsInstalledCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func checkIfWooCommerceIsInstalledParams(cmd *cobra.Command) *client.HostingCheckIfWooCommerceIsInstalledV1Params {
	params := &client.HostingCheckIfWooCommerceIsInstalledV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
