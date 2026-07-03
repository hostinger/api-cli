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

var UpdateVersionCmd = &cobra.Command{
	Use:   "update-version <username> <domain>",
	Short: "Update PHP version",
	Long:  "Changes the PHP version of the website.\n\nUse the Get PHP details endpoint to see the versions available for the website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateVersionBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdatePHPVersionV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateVersionCmd.Flags().StringP("version", "", "", "PHP version to switch the website to.")
	UpdateVersionCmd.MarkFlagRequired("version")
}

func updateVersionBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body
}
