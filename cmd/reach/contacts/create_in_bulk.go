package contacts

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateInBulkCmd = &cobra.Command{
	Use:   "create-in-bulk <profile-uuid>",
	Short: "Create contacts in bulk",
	Long:  "Create many contacts in a profile in a single call.\n\nThe contacts are imported in the background, so a success response means the import was\naccepted rather than finished. Contacts whose email already exists in the profile are\nleft as they are. If double opt-in is enabled, new contacts start off pending and are\nsent a confirmation email.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createInBulkBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateContactsInBulkV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateInBulkCmd.Flags().StringP("contacts", "", "", " (JSON)")
	CreateInBulkCmd.Flags().StringP("note", "", "", "Note applied to every created contact")
	CreateInBulkCmd.Flags().StringSliceP("tag-uuids", "", nil, "Existing tags to attach to every created contact")
	CreateInBulkCmd.MarkFlagRequired("contacts")
}

func createInBulkBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	contactsVal, _ := cmd.Flags().GetString("contacts")
	body["contacts"] = utils.JSONValue(contactsVal, "contacts")
	if cmd.Flags().Changed("note") {
		v, _ := cmd.Flags().GetString("note")
		body["note"] = v
	}
	if cmd.Flags().Changed("tag-uuids") {
		v, _ := cmd.Flags().GetStringSlice("tag-uuids")
		body["tag_uuids"] = v
	}
	return body
}
