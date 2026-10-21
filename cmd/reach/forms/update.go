package forms

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

var UpdateCmd = &cobra.Command{
	Use:   "update <profile-uuid> <form-uuid>",
	Short: "Update form",
	Long:  "Rename a form and/or change its status.\n\nSetting a `draft` or `paused` form to `active` takes it live: its latest saved design is published\nand visitors can submit it. `paused` and `draft` stop the form from accepting submissions. A form is\nnot reachable by visitors until it is taken live. Design changes saved while the form is already\n`active` stay unpublished until the form is paused and set to `active` again. Renaming does not rename\nthe automatic `form:{name}` tag.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "paused", "draft"})
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateFormV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("name", "", "", "Form name. Letters, digits, spaces and - . ' _ only.")
	UpdateCmd.Flags().StringP("status", "", "", "`active` takes the form live, `paused` and `draft` stop it accepting submissions. (one of: active, paused, draft)")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	return body
}
