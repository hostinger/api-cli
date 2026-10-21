package forms

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <profile-uuid>",
	Short: "Create form",
	Long:  "Create an empty draft form.\n\nReach also adds an automatic `form:{name}` tag that every submitter receives. A draft has no design\nand is not reachable by visitors: save a design with the update form template operation, then take\nthe form live by setting its status to `active`. Creating a form does not consume AI credits.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateFormV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "", "", "Form name. Letters, digits, spaces and - . ' _ only.")
	CreateCmd.MarkFlagRequired("name")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
