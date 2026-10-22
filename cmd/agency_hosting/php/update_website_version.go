package php

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateWebsiteVersionCmd = &cobra.Command{
	Use:               "update-website-version <website_uid>",
	Short:             "Update website PHP version",
	Long:              "Switches an Agency Plan website to a different PHP version. Call the available versions endpoint first to see which versions can be selected. The website restarts on the new version, so requests served during the switch may fail and code that is incompatible with the target version will break.",
	Example:           "  hostinger agency-hosting php update-website-version zpwlGlp19 --version 8.2",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateWebsiteVersionBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingUpdateWebsitePHPVersionV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(UpdateWebsiteVersionCmd, utils.StringArg("website_uid"))
	UpdateWebsiteVersionCmd.Flags().StringP("version", "", "", "PHP version to switch the website to, as major.minor. Must be one of the versions returned by the available versions endpoint.")
	UpdateWebsiteVersionCmd.MarkFlagRequired("version")
	UpdateWebsiteVersionCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
}

func updateWebsiteVersionBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body, nil
}
