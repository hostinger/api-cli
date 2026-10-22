package contacts

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateBulkCmd = &cobra.Command{
	Use:               "create-bulk <profile-uuid>",
	Short:             "Create new contacts",
	Long:              "Create a new contact in the email marketing system.\n\nThis endpoint allows you to create a new contact with basic information like name, email, and surname.\n\nIf double opt-in is enabled, the contact will be created with a pending status\nand a confirmation email will be sent.",
	Example:           "  hostinger reach contacts create-bulk 550e8400-e09b-41d4-a716-400055000000 --email john.doe@example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBulkBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateNewContactsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateBulkCmd, utils.StringArg("profile-uuid"))
	CreateBulkCmd.Flags().StringP("email", "", "", "")
	CreateBulkCmd.Flags().StringP("name", "", "", "")
	CreateBulkCmd.Flags().StringP("note", "", "", "")
	CreateBulkCmd.Flags().StringP("phone", "", "", "Phone number in E.164 format (leading \"+\" then 7-15 digits)")
	CreateBulkCmd.Flags().StringP("surname", "", "", "")
	CreateBulkCmd.Flags().StringSliceP("tag-uuids", "", nil, "Existing tags to attach to the created contact")
	CreateBulkCmd.MarkFlagRequired("email")
	CreateBulkCmd.RegisterFlagCompletionFunc("email", cobra.NoFileCompletions)
	CreateBulkCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateBulkCmd.RegisterFlagCompletionFunc("note", cobra.NoFileCompletions)
	CreateBulkCmd.RegisterFlagCompletionFunc("phone", cobra.NoFileCompletions)
	CreateBulkCmd.RegisterFlagCompletionFunc("surname", cobra.NoFileCompletions)
	CreateBulkCmd.RegisterFlagCompletionFunc("tag-uuids", cobra.NoFileCompletions)
}

func createBulkBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	emailVal, _ := cmd.Flags().GetString("email")
	body["email"] = emailVal
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("note") {
		v, _ := cmd.Flags().GetString("note")
		body["note"] = v
	}
	if cmd.Flags().Changed("phone") {
		v, _ := cmd.Flags().GetString("phone")
		body["phone"] = v
	}
	if cmd.Flags().Changed("surname") {
		v, _ := cmd.Flags().GetString("surname")
		body["surname"] = v
	}
	if cmd.Flags().Changed("tag-uuids") {
		v, _ := cmd.Flags().GetStringSlice("tag-uuids")
		body["tag_uuids"] = v
	}
	return body, nil
}
