package contacts

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <profile-uuid> <contact-uuid>",
	Short: "Update a contact",
	Long:  "Update a contact's attributes and custom field values.\n\nOnly the properties present in the request body are changed, so a partial body is enough\nto change a single attribute. Sending a property as `null` clears it.\n\nThe response carries the contact's core attributes. Read back its tags, custom field\nvalues, source and note with `GET /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "subscription-status", []string{"subscribed", "unsubscribed", "confirmed", "pending"})
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateAContactV1WithBodyWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("email", "", "", "")
	UpdateCmd.Flags().StringP("fields", "", "", "Set custom field values. Omit to leave untouched, send an empty array to clear them all. (JSON)")
	UpdateCmd.Flags().StringP("name", "", "", "")
	UpdateCmd.Flags().StringP("note", "", "", "")
	UpdateCmd.Flags().StringP("phone", "", "", "Phone number in E.164 format (leading \"+\" then 7-15 digits)")
	UpdateCmd.Flags().StringP("subscription-status", "", "", "(one of: subscribed, unsubscribed, confirmed, pending)")
	UpdateCmd.Flags().StringP("surname", "", "", "")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("email") {
		v, _ := cmd.Flags().GetString("email")
		body["email"] = v
	}
	if cmd.Flags().Changed("fields") {
		v, _ := cmd.Flags().GetString("fields")
		body["fields"] = utils.JSONValue(v, "fields")
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
	if cmd.Flags().Changed("subscription-status") {
		v, _ := cmd.Flags().GetString("subscription-status")
		body["subscription_status"] = v
	}
	if cmd.Flags().Changed("surname") {
		v, _ := cmd.Flags().GetString("surname")
		body["surname"] = v
	}
	return body
}
