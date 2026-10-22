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

var CreateProfileCmd = &cobra.Command{
	Use:               "create-profile <profile-uuid>",
	Short:             "Create a profile segment",
	Long:              "Create a segment in a profile.\n\nA segment is a saved set of conditions rather than a fixed list, so its membership changes\nas contacts change. Creating one does not modify any contact.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --conditions\n    [\n      {\n        \"attribute\": \"email\",\n        \"operator\": \"equals\",\n        \"value\": \"example@example.com\"\n      }\n    ]",
	Example:           "  hostinger reach segments create-profile 550e8400-e09b-41d4-a716-400055000000 --conditions @conditions.json --logic AND --name 'VIP Customers'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "logic", []string{"AND", "OR"}); err != nil {
			return err
		}
		body, err := createProfileBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateAProfileSegmentV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateProfileCmd, utils.StringArg("profile-uuid"))
	CreateProfileCmd.Flags().StringP("conditions", "", "", "Conditions a contact must satisfy to fall into the segment (JSON, @file or -; shape in the help above)")
	CreateProfileCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions (one of: AND, OR)")
	CreateProfileCmd.Flags().StringP("name", "", "", "")
	CreateProfileCmd.MarkFlagRequired("conditions")
	CreateProfileCmd.MarkFlagRequired("logic")
	CreateProfileCmd.MarkFlagRequired("name")
	CreateProfileCmd.RegisterFlagCompletionFunc("conditions", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateProfileCmd.RegisterFlagCompletionFunc("logic", cobra.FixedCompletions([]cobra.Completion{"AND", "OR"}, cobra.ShellCompDirectiveNoFileComp))
	CreateProfileCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func createProfileBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	conditionsVal, _ := cmd.Flags().GetString("conditions")
	conditionsValJSON, err := utils.JSONValue(conditionsVal, "conditions")
	if err != nil {
		return nil, err
	}
	body["conditions"] = conditionsValJSON
	logicVal, _ := cmd.Flags().GetString("logic")
	body["logic"] = logicVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
