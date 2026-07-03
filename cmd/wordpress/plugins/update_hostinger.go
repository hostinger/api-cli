package plugins

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateHostingerCmd = &cobra.Command{
	Use:   "update-hostinger <username> <software>",
	Short: "Update Hostinger WordPress plugin",
	Long:  "Update a Hostinger plugin to its latest version on a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the update job\nhas been queued.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "slug", []string{"hostinger", "hostinger-ai-assistant", "hostinger-affiliate-plugin", "hostinger-easy-onboarding", "hostinger-reach"})
		payload, err := json.Marshal(updateHostingerBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdateHostingerWordPressPluginV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateHostingerCmd.Flags().StringP("slug", "", "", "Slug of the Hostinger plugin to update to its latest version. (one of: hostinger, hostinger-ai-assistant, hostinger-affiliate-plugin, hostinger-easy-onboarding, hostinger-reach)")
	UpdateHostingerCmd.MarkFlagRequired("slug")
}

func updateHostingerBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	return body
}
