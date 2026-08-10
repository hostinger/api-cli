package contact_fields

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

var CreateCmd = &cobra.Command{
	Use:   "create <profile-uuid>",
	Short: "Create a contact field",
	Long:  "Define a new custom contact field in a profile.\n\nThe `slug` is derived from the label and, like the field type, cannot be changed later.\nUse the returned uuid to set values on contacts.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"text", "number", "date", "single_choice", "multi_choice"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateAContactFieldV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("label", "", "", "")
	CreateCmd.Flags().StringSliceP("options", "", nil, "Required for single_choice and multi_choice, ignored for the scalar types. Labels must be unique regardless of casing.")
	CreateCmd.Flags().StringP("type", "", "", "Immutable once the field exists (one of: text, number, date, single_choice, multi_choice)")
	CreateCmd.MarkFlagRequired("label")
	CreateCmd.MarkFlagRequired("type")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	labelVal, _ := cmd.Flags().GetString("label")
	body["label"] = labelVal
	if cmd.Flags().Changed("options") {
		v, _ := cmd.Flags().GetStringSlice("options")
		body["options"] = v
	}
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	return body
}
