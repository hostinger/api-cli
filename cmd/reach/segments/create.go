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

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create a new contact segment",
	Long:              "Create a new contact segment.\n\nThis endpoint allows creating a new contact segment that can be used to organize contacts.\nThe segment can be configured with specific criteria like email, name, subscription status, etc.\n\nDeprecated. This endpoint cannot target a profile, so it always falls back to\nthe client's default profile and cannot create segments in any other profile. Use\n`POST /api/reach/v1/profiles/{profileUuid}/segmentation/segments` instead.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --conditions\n    [\n      {\n        \"attribute\": \"note\",\n        \"operator\": \"equals\",\n        \"value\": \"string\"\n      }\n    ]",
	Example:           "  hostinger reach segments create --conditions '[{\"attribute\":\"note\",\"operator\":\"equals\",\"value\":\"string\"}]' --logic AND --name 'New segment name'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "logic", []string{"AND", "OR"}); err != nil {
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
			r, err := api.Request().ReachCreateANewContactSegmentV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("conditions", "", "", "(JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("logic", "", "", "(one of: AND, OR)")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.MarkFlagRequired("conditions")
	CreateCmd.MarkFlagRequired("logic")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.RegisterFlagCompletionFunc("conditions", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("logic", cobra.FixedCompletions([]cobra.Completion{"AND", "OR"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
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
