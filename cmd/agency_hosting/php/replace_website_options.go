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

var ReplaceWebsiteOptionsCmd = &cobra.Command{
	Use:               "replace-website-options <website_uid>",
	Short:             "Replace website PHP options",
	Long:              "Replaces the custom php.ini values on an Agency Plan website with the ones provided. Any option not in the request is reset to its default, so call the options endpoint first and send the full desired set. Sending an empty array resets every option to its default.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --options\n    [\n      {\n        \"name\": \"memory_limit\",\n        \"value\": \"256M\"\n      }\n    ]",
	Example:           "  hostinger agency-hosting php replace-website-options zpwlGlp19 --options '[{\"name\":\"memory_limit\",\"value\":\"256M\"}]'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := replaceWebsiteOptionsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingReplaceWebsitePHPOptionsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ReplaceWebsiteOptionsCmd, utils.StringArg("website_uid"))
	ReplaceWebsiteOptionsCmd.Flags().StringP("options", "", "", "Option names and values. Each name must be one of the options returned by the options endpoint, and each value must satisfy that option's allowed_values when it declares them. (JSON, @file or -; shape in the help above)")
	ReplaceWebsiteOptionsCmd.MarkFlagRequired("options")
	ReplaceWebsiteOptionsCmd.RegisterFlagCompletionFunc("options", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func replaceWebsiteOptionsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	optionsVal, _ := cmd.Flags().GetString("options")
	optionsValJSON, err := utils.JSONValue(optionsVal, "options")
	if err != nil {
		return nil, err
	}
	body["options"] = optionsValJSON
	return body, nil
}
