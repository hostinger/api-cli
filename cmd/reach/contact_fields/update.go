package contact_fields

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
	Use:   "update <profile-uuid> <field-uuid>",
	Short: "Update a contact field",
	Long:  "Rename a custom contact field and, for the choice types, replace its option set.\n\nOptions carrying a uuid are kept and relabelled, options without one are created, and any\nexisting option left out of the list is deleted along with the values contacts hold for\nit. The field type and slug cannot be changed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateAContactFieldV1WithBodyWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("label", "", "", "")
	UpdateCmd.Flags().StringP("options", "", "", "Replaces the option set when provided. Entries carrying a uuid are kept and relabelled, entries without one are created, and any existing option missing from the list is deleted along with the values contacts hold for it. (JSON)")
	UpdateCmd.MarkFlagRequired("label")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	labelVal, _ := cmd.Flags().GetString("label")
	body["label"] = labelVal
	if cmd.Flags().Changed("options") {
		v, _ := cmd.Flags().GetString("options")
		body["options"] = utils.JSONValue(v, "options")
	}
	return body
}
