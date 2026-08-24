package segments

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

var PreviewContactsMatchingConditionsCmd = &cobra.Command{
	Use:   "preview-contacts-matching-conditions <profile-uuid>",
	Short: "Preview contacts matching conditions",
	Long:  "Preview the contacts matching a set of conditions without saving a segment.\n\nThe body is the same set of conditions accepted when creating or updating a segment, so this\nis how to check who a filter reaches, and how many, before persisting it. Nothing is stored\nand no contact is modified.\n\nCall the segment filter attributes endpoint first to discover the valid `attribute`,\n`operator` and `value` combinations.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "logic", []string{"AND", "OR"})
		utils.EnumCheck(cmd, "sort-by", []string{"email", "name", "surname", "phone", "subscription_status"})
		utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"})
		payload, err := json.Marshal(previewContactsMatchingConditionsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachPreviewContactsMatchingConditionsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PreviewContactsMatchingConditionsCmd.Flags().StringP("conditions", "", "", "Conditions a contact must satisfy to appear in the preview (JSON)")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions (one of: AND, OR)")
	PreviewContactsMatchingConditionsCmd.Flags().IntP("page", "", 0, "Page number")
	PreviewContactsMatchingConditionsCmd.Flags().IntP("per-page", "", 0, "Number of items per page")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("search", "", "", "Narrow the preview to contacts whose email matches")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("sort-by", "", "", "(one of: email, name, surname, phone, subscription_status)")
	PreviewContactsMatchingConditionsCmd.Flags().StringP("sort-direction", "", "", "(one of: asc, desc)")
	PreviewContactsMatchingConditionsCmd.MarkFlagRequired("conditions")
	PreviewContactsMatchingConditionsCmd.MarkFlagRequired("logic")
}

func previewContactsMatchingConditionsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	conditionsVal, _ := cmd.Flags().GetString("conditions")
	body["conditions"] = utils.JSONValue(conditionsVal, "conditions")
	logicVal, _ := cmd.Flags().GetString("logic")
	body["logic"] = logicVal
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		body["page"] = v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		body["per_page"] = v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		body["search"] = v
	}
	if cmd.Flags().Changed("sort-by") {
		v, _ := cmd.Flags().GetString("sort-by")
		body["sort_by"] = v
	}
	if cmd.Flags().Changed("sort-direction") {
		v, _ := cmd.Flags().GetString("sort-direction")
		body["sort_direction"] = v
	}
	return body
}
