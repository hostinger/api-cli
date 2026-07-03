package plugins

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CheckIfWooCommerceIsInstalledCmd = &cobra.Command{
	Use:   "check-if-woo-commerce-is-installed",
	Short: "Check if WooCommerce is installed",
	Long:  "Check whether WooCommerce is installed on any WordPress installation of a\ndomain. Optionally filter by domain to scope the check.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingCheckIfWooCommerceIsInstalledV1WithResponse(context.TODO(), checkIfWooCommerceIsInstalledParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CheckIfWooCommerceIsInstalledCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
}

func checkIfWooCommerceIsInstalledParams(cmd *cobra.Command) *client.HostingCheckIfWooCommerceIsInstalledV1Params {
	params := &client.HostingCheckIfWooCommerceIsInstalledV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
