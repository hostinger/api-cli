package wordpress

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ChangePlanWebsiteCoreVersionCmd = &cobra.Command{
	Use:   "change-plan-website-core-version <website_uid>",
	Short: "Change Agency Plan website WordPress core version",
	Long:  "Changes the installed WordPress core version on an Agency Plan website to one of the versions available for installation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changePlanWebsiteCoreVersionBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingChangeAgencyPlanWebsiteWordPressCoreVersionV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangePlanWebsiteCoreVersionCmd.Flags().StringP("version", "", "", "Target WordPress core version to install. Must be one of the available versions.")
	ChangePlanWebsiteCoreVersionCmd.MarkFlagRequired("version")
}

func changePlanWebsiteCoreVersionBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body
}
