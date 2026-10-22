package tags

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:               "rename <profile-uuid> <tag-uuid>",
	Short:             "Rename a tag",
	Long:              "Rename a tag.\n\nThe contacts assigned to the tag are unaffected. Names are unique within a profile, so\nrenaming a tag to a name that is already taken is rejected.",
	Example:           "  hostinger reach tags rename 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 --value Newsletter",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := renameBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachRenameATagV1WithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RenameCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"))
	RenameCmd.Flags().StringP("value", "", "", "New tag name")
	RenameCmd.MarkFlagRequired("value")
	RenameCmd.RegisterFlagCompletionFunc("value", cobra.NoFileCompletions)
}

func renameBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	valueVal, _ := cmd.Flags().GetString("value")
	body["value"] = valueVal
	return body, nil
}
