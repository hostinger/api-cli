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

var UpdateExtensionsCmd = &cobra.Command{
	Use:   "update-extensions <username> <domain>",
	Short: "Update PHP extensions",
	Long:  "Enables or disables PHP extensions (modules) for the website.\n\nUse the Get PHP details endpoint to check the current extension states before changing them.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateExtensionsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdatePHPExtensionsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateExtensionsCmd.Flags().StringSliceP("disable", "", nil, "PHP extensions to disable.")
	UpdateExtensionsCmd.Flags().StringSliceP("enable", "", nil, "PHP extensions to enable.")
}

func updateExtensionsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("disable") {
		v, _ := cmd.Flags().GetStringSlice("disable")
		body["disable"] = v
	}
	if cmd.Flags().Changed("enable") {
		v, _ := cmd.Flags().GetStringSlice("enable")
		body["enable"] = v
	}
	return body
}
