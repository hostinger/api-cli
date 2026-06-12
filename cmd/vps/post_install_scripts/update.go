package post_install_scripts

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

var UpdateCmd = &cobra.Command{
	Use:   "update <post-install-script-id>",
	Short: "Update post-install script",
	Long:  "Update a specific post-install script.\n\nUse this endpoint to modify existing automation scripts.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSUpdatePostInstallScriptV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("content", "", "", "Content of the script")
	UpdateCmd.Flags().StringP("name", "", "", "Name of the script")
	UpdateCmd.MarkFlagRequired("content")
	UpdateCmd.MarkFlagRequired("name")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
