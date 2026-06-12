package contacts

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateBulkCmd = &cobra.Command{
	Use:   "create-bulk <profile-uuid>",
	Short: "Create new contacts",
	Long:  "Create a new contact in the email marketing system.\n\nThis endpoint allows you to create a new contact with basic information like name, email, and surname.\n\nIf double opt-in is enabled, the contact will be created with a pending status\nand a confirmation email will be sent.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBulkBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateNewContactsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateBulkCmd.Flags().StringP("email", "", "", "")
	CreateBulkCmd.Flags().StringP("name", "", "", "")
	CreateBulkCmd.Flags().StringP("note", "", "", "")
	CreateBulkCmd.Flags().StringP("phone", "", "", "Phone number in E.164 format (leading \"+\" then 7-15 digits)")
	CreateBulkCmd.Flags().StringP("surname", "", "", "")
	CreateBulkCmd.MarkFlagRequired("email")
}

func createBulkBody(cmd *cobra.Command) map[string]any {
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
	return body
}
