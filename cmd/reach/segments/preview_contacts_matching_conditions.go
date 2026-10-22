package segments

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PreviewContactsMatchingConditionsCmd = &cobra.Command{
	Use:               "preview-contacts-matching-conditions <profile-uuid>",
	Short:             "Preview contacts matching conditions",
	Long:              "Preview the contacts matching a set of conditions without saving a segment.\n\nThe body is the same set of conditions accepted when creating or updating a segment, so this\nis how to check who a filter reaches, and how many, before persisting it. Nothing is stored\nand no contact is modified.\n\nCall the segment filter attributes endpoint first to discover the valid `attribute`,\n`operator` and `value` combinations.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --conditions\n    [\n      {\n        \"attribute\": \"email\",\n        \"operator\": \"equals\",\n        \"value\": \"example@example.com\"\n      }\n    ]",
	Example:           "  hostinger reach segments preview-contacts-matching-conditions 550e8400-e09b-41d4-a716-400055000000 --conditions @conditions.json --logic AND",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "logic", []string{"AND", "OR"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "sort-by", []string{"email", "name", "surname", "phone", "subscription_status"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"}); err != nil {
			return err
		}
		body, err := previewContactsMatchingConditionsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachPreviewContactsMatchingConditionsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(PreviewContactsMatchingConditionsCmd, utils.StringArg("profile-uuid"))
	PreviewContactsMatchingConditionsCmd.Flags().StringP("conditions", "", "", "Conditions a contact must satisfy to appear in the preview (JSON, @file or -; shape in the help above)")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions (one of: AND, OR)")
	PreviewContactsMatchingConditionsCmd.Flags().IntP("page", "", 0, "Page number")
	PreviewContactsMatchingConditionsCmd.Flags().IntP("per-page", "", 0, "Number of items per page")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("search", "", "", "Narrow the preview to contacts whose email matches")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("sort-by", "", "", "(one of: email, name, surname, phone, subscription_status)")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("sort-direction", "", "", "(one of: asc, desc)")
	PreviewContactsMatchingConditionsCmd.MarkFlagRequired("conditions")
	PreviewContactsMatchingConditionsCmd.MarkFlagRequired("logic")
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("conditions", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("logic", cobra.FixedCompletions([]cobra.Completion{"AND", "OR"}, cobra.ShellCompDirectiveNoFileComp))
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("sort-by", cobra.FixedCompletions([]cobra.Completion{"email", "name", "surname", "phone", "subscription_status"}, cobra.ShellCompDirectiveNoFileComp))
	PreviewContactsMatchingConditionsCmd.RegisterFlagCompletionFunc("sort-direction", cobra.FixedCompletions([]cobra.Completion{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
}

func previewContactsMatchingConditionsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	conditionsVal, _ := cmd.Flags().GetString("conditions")
	conditionsValJSON, err := utils.JSONValue(conditionsVal, "conditions")
	if err != nil {
		return nil, err
	}
	body["conditions"] = conditionsValJSON
	logicVal, _ := cmd.Flags().GetString("logic")
	body["logic"] = logicVal
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		body["page"] = v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		body["per_page"] = v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		body["search"] = v
	}
	if cmd.Flags().Changed("sort-by") {
		v, _ := cmd.Flags().GetString("sort-by")
		body["sort_by"] = v
	}
	if cmd.Flags().Changed("sort-direction") {
		v, _ := cmd.Flags().GetString("sort-direction")
		body["sort_direction"] = v
	}
	return body, nil
}
