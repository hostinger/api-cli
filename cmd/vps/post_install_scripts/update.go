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

var UpdateCmd = &cobra.Command{
	Use:     "update <post-install-script-id>",
	Short:   "Update post-install script",
	Long:    "Update a specific post-install script.\n\nUse this endpoint to modify existing automation scripts.",
	Example: "  hostinger vps post-install-scripts update 9568314 --content '#!/bin/bash\n\necho '\\''Hello, World!'\\''' --name 'My Script'",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetPostInstallScriptsV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSUpdatePostInstallScriptV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.IntArg("post-install-script-id"))
	UpdateCmd.Flags().StringP("content", "", "", "Content of the script")
	UpdateCmd.Flags().StringP("name", "", "", "Name of the script")
	UpdateCmd.MarkFlagRequired("content")
	UpdateCmd.MarkFlagRequired("name")
	UpdateCmd.RegisterFlagCompletionFunc("content", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
