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

var UpdateWebsiteVersionCmd = &cobra.Command{
	Use:   "update-website-version <website_uid>",
	Short: "Update website PHP version",
	Long:  "Switches an Agency Plan website to a different PHP version. Call the available versions endpoint first to see which versions can be selected. The website restarts on the new version, so requests served during the switch may fail and code that is incompatible with the target version will break.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateWebsiteVersionBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingUpdateWebsitePHPVersionV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateWebsiteVersionCmd.Flags().StringP("version", "", "", "PHP version to switch the website to, as major.minor. Must be one of the versions returned by the available versions endpoint.")
	UpdateWebsiteVersionCmd.MarkFlagRequired("version")
}

func updateWebsiteVersionBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body
}
