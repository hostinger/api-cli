package plugins

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListInstalledCmd = &cobra.Command{
	Use:   "list-installed <username> <software>",
	Short: "List installed WordPress plugins",
	Long:  "List plugins installed on a WordPress installation, including their status,\navailable updates and known vulnerabilities.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "category", []string{"cache"})
		r, err := api.Request().HostingListInstalledWordPressPluginsV1WithResponse(context.TODO(), args[0], args[1], listInstalledParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListInstalledCmd.Flags().StringP("category", "", "", "Filter installed plugins by category. (one of: cache)")
}

func listInstalledParams(cmd *cobra.Command) *client.HostingListInstalledWordPressPluginsV1Params {
	params := &client.HostingListInstalledWordPressPluginsV1Params{}
	if cmd.Flags().Changed("category") {
		v, _ := cmd.Flags().GetString("category")
		e := client.HostingListInstalledWordPressPluginsV1ParamsCategory(v)
		params.Category = &e
	}
	return params
}
