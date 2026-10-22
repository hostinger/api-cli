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

var ReplaceWebsiteExtensionsCmd = &cobra.Command{
	Use:               "replace-website-extensions <website_uid>",
	Short:             "Replace website PHP extensions",
	Long:              "Replaces the set of PHP extensions enabled on an Agency Plan website with the ones provided. Any toggleable extension not in the request is disabled, so call the extensions endpoint first and send the full desired set. Extensions compiled into PHP, reported with the \"built-in\" state, are always active and are unaffected.",
	Example:           "  hostinger agency-hosting php replace-website-extensions zpwlGlp19 --extensions bz2,zip",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := replaceWebsiteExtensionsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingReplaceWebsitePHPExtensionsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ReplaceWebsiteExtensionsCmd, utils.StringArg("website_uid"))
	ReplaceWebsiteExtensionsCmd.Flags().StringSliceP("extensions", "", nil, "Extension names, exactly as returned by the extensions endpoint.")
	ReplaceWebsiteExtensionsCmd.MarkFlagRequired("extensions")
	ReplaceWebsiteExtensionsCmd.RegisterFlagCompletionFunc("extensions", cobra.NoFileCompletions)
}

func replaceWebsiteExtensionsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	extensionsVal, _ := cmd.Flags().GetStringSlice("extensions")
	body["extensions"] = extensionsVal
	return body, nil
}
