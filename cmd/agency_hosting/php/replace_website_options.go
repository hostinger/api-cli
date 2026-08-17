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

var ReplaceWebsiteOptionsCmd = &cobra.Command{
	Use:   "replace-website-options <website_uid>",
	Short: "Replace website PHP options",
	Long:  "Replaces the custom php.ini values on an Agency Plan website with the ones provided. Any option not in the request is reset to its default, so call the options endpoint first and send the full desired set. Sending an empty array resets every option to its default.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(replaceWebsiteOptionsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingReplaceWebsitePHPOptionsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ReplaceWebsiteOptionsCmd.Flags().StringP("options", "", "", "Option names and values. Each name must be one of the options returned by the options endpoint, and each value must satisfy that option's allowed_values when it declares them. (JSON)")
	ReplaceWebsiteOptionsCmd.MarkFlagRequired("options")
}

func replaceWebsiteOptionsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	optionsVal, _ := cmd.Flags().GetString("options")
	body["options"] = utils.JSONValue(optionsVal, "options")
	return body
}
