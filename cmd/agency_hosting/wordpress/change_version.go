package wordpress

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ChangeVersionCmd = &cobra.Command{
	Use:               "change-version <website_uid>",
	Short:             "Change WordPress version",
	Long:              "Changes the installed WordPress core version on an Agency Plan website to one of the versions available for installation.",
	Example:           "  hostinger agency-hosting wordpress change-version zpwlGlp19 --version 6.5.5",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := changeVersionBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingChangeWordPressVersionV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ChangeVersionCmd, utils.StringArg("website_uid"))
	ChangeVersionCmd.Flags().StringP("version", "", "", "Target WordPress core version to install. Must be one of the available versions.")
	ChangeVersionCmd.MarkFlagRequired("version")
	ChangeVersionCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
}

func changeVersionBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body, nil
}
