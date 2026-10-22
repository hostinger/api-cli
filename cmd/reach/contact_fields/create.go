package contact_fields

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
	Short:             "Create a contact field",
	Long:              "Define a new custom contact field in a profile.\n\nThe `slug` is derived from the label and, like the field type, cannot be changed later.\nUse the returned uuid to set values on contacts.",
	Example:           "  hostinger reach contact-fields create 550e8400-e09b-41d4-a716-400055000000 --label 'Job title' --type text",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "type", []string{"text", "number", "date", "single_choice", "multi_choice"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateAContactFieldV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("label", "", "", "")
	CreateCmd.Flags().StringSliceP("options", "", nil, "Required for single_choice and multi_choice, ignored for the scalar types. Labels must be unique regardless of casing.")
	CreateCmd.Flags().StringP("type", "", "", "Immutable once the field exists (one of: text, number, date, single_choice, multi_choice)")
	CreateCmd.MarkFlagRequired("label")
	CreateCmd.MarkFlagRequired("type")
	CreateCmd.RegisterFlagCompletionFunc("label", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("options", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"text", "number", "date", "single_choice", "multi_choice"}, cobra.ShellCompDirectiveNoFileComp))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	labelVal, _ := cmd.Flags().GetString("label")
	body["label"] = labelVal
	if cmd.Flags().Changed("options") {
		v, _ := cmd.Flags().GetStringSlice("options")
		body["options"] = v
	}
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	return body, nil
}
