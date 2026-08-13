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

var ChangeVersionCmd = &cobra.Command{
	Use:   "change-version <website_uid>",
	Short: "Change WordPress version",
	Long:  "Changes the installed WordPress core version on an Agency Plan website to one of the versions available for installation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changeVersionBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingChangeWordPressVersionV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangeVersionCmd.Flags().StringP("version", "", "", "Target WordPress core version to install. Must be one of the available versions.")
	ChangeVersionCmd.MarkFlagRequired("version")
}

func changeVersionBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body
}
