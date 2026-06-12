package post_install_scripts

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create post-install script",
	Long:  "Add a new post-install script to your account, which can then be used after virtual machine installation.\n\nThe script contents will be saved to the file `/post_install` with executable attribute set\nand will be executed once virtual machine is installed.\nThe output of the script will be redirected to `/post_install.log`. Maximum script size is 48KB.\n\nUse this endpoint to create automation scripts for VPS setup tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSCreatePostInstallScriptV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("content", "", "", "Content of the script")
	CreateCmd.Flags().StringP("name", "", "", "Name of the script")
	CreateCmd.MarkFlagRequired("content")
	CreateCmd.MarkFlagRequired("name")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
