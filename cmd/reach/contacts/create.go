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

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create a new contact",
	Long:              "Create a new contact in the email marketing system.\n\nThis endpoint allows you to create a new contact with basic information like name, email, and surname.\n\nIf double opt-in is enabled,\nthe contact will be created with a pending status and a confirmation email will be sent.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --fields\n    [\n      {\n        \"selected_option_uuids\": [\n          \"string\"\n        ],\n        \"uuid\": \"string\",\n        \"value\": \"string\"\n      }\n    ]",
	Example:           "  hostinger reach contacts create --email john.doe@example.com",
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
			r, err := api.Request().ReachCreateANewContactV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("email", "", "", "")
	CreateCmd.Flags().StringP("fields", "", "", "Custom field values to set on the contact. When a contact with this email already exists, only fields that have no value yet are filled. Unknown field UUIDs and values that do not match the field type are skipped. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.Flags().StringP("note", "", "", "")
	CreateCmd.Flags().StringP("phone", "", "", "Phone number in E.164 format (leading \"+\" then 7-15 digits)")
	CreateCmd.Flags().StringP("surname", "", "", "")
	CreateCmd.Flags().StringSliceP("tag-uuids", "", nil, "Existing tags to attach to the created contact")
	CreateCmd.MarkFlagRequired("email")
	CreateCmd.RegisterFlagCompletionFunc("email", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("fields", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("note", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("phone", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("surname", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("tag-uuids", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	emailVal, _ := cmd.Flags().GetString("email")
	body["email"] = emailVal
	if cmd.Flags().Changed("fields") {
		fieldsVal, _ := cmd.Flags().GetString("fields")
		fieldsValJSON, err := utils.JSONValue(fieldsVal, "fields")
		if err != nil {
			return nil, err
		}
		body["fields"] = fieldsValJSON
	}
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
