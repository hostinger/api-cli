package php

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

var UpdateOptionsCmd = &cobra.Command{
	Use:   "update-options <username> <domain>",
	Short: "Update PHP options",
	Long:  "Updates PHP options for the website (e.g. `memory_limit`, `max_execution_time`, `upload_max_filesize`).\nOnly provide the options you want to change, inside the `options` object.\n\nValues above the account plan limit are silently capped to that limit, so the request can succeed\nwith a smaller applied value. Call the Get PHP details endpoint afterwards to read the applied value.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateOptionsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdatePHPOptionsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateOptionsCmd.Flags().StringP("options", "", "", "Map of PHP options to update, keyed by option name. Only include options you want to change. (JSON)")
	UpdateOptionsCmd.MarkFlagRequired("options")
}

func updateOptionsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	optionsVal, _ := cmd.Flags().GetString("options")
	body["options"] = utils.JSONValue(optionsVal, "options")
	return body
}
