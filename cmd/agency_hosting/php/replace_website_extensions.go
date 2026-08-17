package php

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ReplaceWebsiteExtensionsCmd = &cobra.Command{
	Use:   "replace-website-extensions <website_uid>",
	Short: "Replace website PHP extensions",
	Long:  "Replaces the set of PHP extensions enabled on an Agency Plan website with the ones provided. Any toggleable extension not in the request is disabled, so call the extensions endpoint first and send the full desired set. Extensions compiled into PHP, reported with the \"built-in\" state, are always active and are unaffected.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(replaceWebsiteExtensionsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingReplaceWebsitePHPExtensionsV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ReplaceWebsiteExtensionsCmd.Flags().StringSliceP("extensions", "", nil, "Extension names, exactly as returned by the extensions endpoint.")
	ReplaceWebsiteExtensionsCmd.MarkFlagRequired("extensions")
}

func replaceWebsiteExtensionsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	extensionsVal, _ := cmd.Flags().GetStringSlice("extensions")
	body["extensions"] = extensionsVal
	return body
}
