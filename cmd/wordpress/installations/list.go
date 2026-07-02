package installations

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
	Use:   "list",
	Short: "List WordPress installations",
	Long:  "List WordPress installations accessible to the authenticated client.\n\nUse this endpoint to discover existing WordPress installations and to poll\nfor installation status after calling the install endpoint. When a newly\nrequested installation appears in this list, WordPress is ready. Filter by\nusername and domain to narrow results to a specific website.\n\nEach installation includes a `valid` flag and, when invalid, a\n`validationError` describing why.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "ownership", []string{"owned", "managed", "all"})
		r, err := api.Request().HostingListWordPressInstallationsV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("username", "", "", "Filter by specific username")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
	ListCmd.Flags().StringP("ownership", "", "", "Filter by ownership type. Defaults to \"owned\". Use \"all\" to include both owned and managed installations. (one of: owned, managed, all)")
}

func listParams(cmd *cobra.Command) *client.HostingListWordPressInstallationsV1Params {
	params := &client.HostingListWordPressInstallationsV1Params{}
	if cmd.Flags().Changed("username") {
		v, _ := cmd.Flags().GetString("username")
		params.Username = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("ownership") {
		v, _ := cmd.Flags().GetString("ownership")
		e := client.HostingListWordPressInstallationsV1ParamsOwnership(v)
		params.Ownership = &e
	}
	return params
}
