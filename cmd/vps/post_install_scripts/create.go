package post_install_scripts

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create post-install script",
	Long:              "Add a new post-install script to your account, which can then be used after virtual machine installation.\n\nThe script contents will be saved to the file `/post_install` with executable attribute set\nand will be executed once virtual machine is installed.\nThe output of the script will be redirected to `/post_install.log`. Maximum script size is 48KB.\n\nUse this endpoint to create automation scripts for VPS setup tasks.",
	Example:           "  hostinger vps post-install-scripts create --content '#!/bin/bash\n\necho '\\''Hello, World!'\\''' --name 'My Script'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSCreatePostInstallScriptV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd)
	CreateCmd.Flags().StringP("content", "", "", "Content of the script")
	CreateCmd.Flags().StringP("name", "", "", "Name of the script")
	CreateCmd.MarkFlagRequired("content")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.RegisterFlagCompletionFunc("content", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
