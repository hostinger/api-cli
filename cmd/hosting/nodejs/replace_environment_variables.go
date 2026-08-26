package nodejs

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

var ReplaceEnvironmentVariablesCmd = &cobra.Command{
	Use:   "replace-environment-variables <username> <domain>",
	Short: "Replace Node.js environment variables",
	Long:  "Replaces the website's Node.js environment variables with the ones provided. This is a\nfull replace: any variable not in the request is deleted, and sending an empty `env_vars`\narray deletes every variable. Saving writes the values and restarts the running Node.js\nprocess.\n\nA restart is enough for apps that read environment variables at process start, such as\nExpress or NestJS. It is not enough for frameworks that bake variables into the build.\nNext.js standalone is one of those: build-time values (including `NEXT_PUBLIC_*`) need a\nfresh build. After this call, use the `Start Node.js build` endpoint so those apps\npick up the new values.\n\nThe `List Node.js environment variables` endpoint returns masked values (`********`), so\nnever copy values from it into this request. Always send the full desired set with real\nvalues taken from the project `.env` file or the user prompt.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(replaceEnvironmentVariablesBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingReplaceNodeJsEnvironmentVariablesV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ReplaceEnvironmentVariablesCmd.Flags().StringP("env-vars", "", "", "Environment variables to set. This is the full desired set: any variable not in\nthis list is deleted, and an empty array deletes every variable. (JSON)")
	ReplaceEnvironmentVariablesCmd.MarkFlagRequired("env-vars")
}

func replaceEnvironmentVariablesBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	envVarsVal, _ := cmd.Flags().GetString("env-vars")
	body["env_vars"] = utils.JSONValue(envVarsVal, "env-vars")
	return body
}
