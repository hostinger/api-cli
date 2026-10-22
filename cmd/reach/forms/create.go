package forms

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
	Use:               "create <profile-uuid>",
	Short:             "Create form",
	Long:              "Create an empty draft form.\n\nReach also adds an automatic `form:{name}` tag that every submitter receives. A draft has no design\nand is not reachable by visitors: save a design with the update form template operation, then take\nthe form live by setting its status to `active`. Creating a form does not consume AI credits.",
	Example:           "  hostinger reach forms create 550e8400-e09b-41d4-a716-400055000000 --name 'Newsletter signup'",
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
			r, err := api.Request().ReachCreateFormV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("profile-uuid"))
	CreateCmd.Flags().StringP("name", "", "", "Form name. Letters, digits, spaces and - . ' _ only.")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
