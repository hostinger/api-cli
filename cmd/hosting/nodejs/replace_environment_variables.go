package nodejs

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ReplaceEnvironmentVariablesCmd = &cobra.Command{
	Use:     "replace-environment-variables <username> <domain>",
	Short:   "Replace Node.js environment variables",
	Long:    "Replaces the website's Node.js environment variables with the ones provided. This is a\nfull replace: any variable not in the request is deleted, and sending an empty `env_vars`\narray deletes every variable. Saving writes the values and restarts the running Node.js\nprocess.\n\nA restart is enough for apps that read environment variables at process start, such as\nExpress or NestJS. It is not enough for frameworks that bake variables into the build.\nNext.js standalone is one of those: build-time values (including `NEXT_PUBLIC_*`) need a\nfresh build. After this call, use the `Start Node.js build` endpoint so those apps\npick up the new values.\n\nThe `List Node.js environment variables` endpoint returns masked values (`********`), so\nnever copy values from it into this request. Always send the full desired set with real\nvalues taken from the project `.env` file or the user prompt.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --env-vars\n    [\n      {\n        \"key\": \"API_URL\",\n        \"value\": \"https://api.example.com\"\n      }\n    ]",
	Example: "  hostinger hosting nodejs replace-environment-variables u123456789 mydomain.tld --env-vars '[{\"key\":\"API_URL\",\"value\":\"https://api.example.com\"}]'",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := replaceEnvironmentVariablesBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingReplaceNodeJsEnvironmentVariablesV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ReplaceEnvironmentVariablesCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ReplaceEnvironmentVariablesCmd.Flags().StringP("env-vars", "", "", "Environment variables to set. This is the full desired set: any variable not in this list is deleted, and an empty array deletes every variable. (JSON, @file or -; shape in the help above)")
	ReplaceEnvironmentVariablesCmd.MarkFlagRequired("env-vars")
	ReplaceEnvironmentVariablesCmd.RegisterFlagCompletionFunc("env-vars", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func replaceEnvironmentVariablesBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	envVarsVal, _ := cmd.Flags().GetString("env-vars")
	envVarsValJSON, err := utils.JSONValue(envVarsVal, "env-vars")
	if err != nil {
		return nil, err
	}
	body["env_vars"] = envVarsValJSON
	return body, nil
}
