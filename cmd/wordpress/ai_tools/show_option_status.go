package ai_tools

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ShowOptionStatusCmd = &cobra.Command{
	Use:   "show-option-status <username> <software>",
	Short: "Show AI option status",
	Long:  "Show the current AI option status for the Hostinger Tools plugin on the\nspecified WordPress installation. Filter by `option` to return a single\noption, or omit it to return all options.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "option", []string{"llmstxt", "web2agent"})
		r, err := api.Request().HostingShowAIOptionStatusV1WithResponse(context.TODO(), args[0], args[1], showOptionStatusParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ShowOptionStatusCmd.Flags().StringP("option", "", "", "Filter the status by a single AI option. (one of: llmstxt, web2agent)")
}

func showOptionStatusParams(cmd *cobra.Command) *client.HostingShowAIOptionStatusV1Params {
	params := &client.HostingShowAIOptionStatusV1Params{}
	if cmd.Flags().Changed("option") {
		v, _ := cmd.Flags().GetString("option")
		e := client.HostingShowAIOptionStatusV1ParamsOption(v)
		params.Option = &e
	}
	return params
}
