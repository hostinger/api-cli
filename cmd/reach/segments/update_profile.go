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

var UpdateProfileCmd = &cobra.Command{
	Use:     "update-profile <profile-uuid> <segment-uuid>",
	Short:   "Update a profile segment",
	Long:    "Rename a segment and/or replace the conditions that define it.\n\n`name` is always required. Omit `conditions` to rename without touching the conditions;\nsupply them and they replace the existing set entirely rather than being merged into it.\nContacts are never modified, but which of them match the segment can change immediately.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --conditions\n    [\n      {\n        \"attribute\": \"email\",\n        \"operator\": \"equals\",\n        \"value\": \"example@example.com\"\n      }\n    ]",
	Example: "  hostinger reach segments update-profile 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --name 'VIP Customers'",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListSegmentsV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "logic", []string{"AND", "OR"}); err != nil {
			return err
		}
		body, err := updateProfileBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachUpdateAProfileSegmentV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateProfileCmd, utils.StringArg("profile-uuid"), utils.StringArg("segment-uuid"))
	UpdateProfileCmd.Flags().StringP("conditions", "", "", "Replaces the existing conditions entirely. Omit to keep the current ones. (JSON, @file or -; shape in the help above)")
	UpdateProfileCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions. Required when conditions are given. (one of: AND, OR)")
	UpdateProfileCmd.Flags().StringP("name", "", "", "")
	UpdateProfileCmd.MarkFlagRequired("name")
	UpdateProfileCmd.RegisterFlagCompletionFunc("conditions", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateProfileCmd.RegisterFlagCompletionFunc("logic", cobra.FixedCompletions([]cobra.Completion{"AND", "OR"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateProfileCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func updateProfileBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("conditions") {
		conditionsVal, _ := cmd.Flags().GetString("conditions")
		conditionsValJSON, err := utils.JSONValue(conditionsVal, "conditions")
		if err != nil {
			return nil, err
		}
		body["conditions"] = conditionsValJSON
	}
	if cmd.Flags().Changed("logic") {
		v, _ := cmd.Flags().GetString("logic")
		body["logic"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
