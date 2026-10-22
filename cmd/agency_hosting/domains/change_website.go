package domains

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ChangeWebsiteCmd = &cobra.Command{
	Use:               "change-website <website_uid> <from_domain>",
	Short:             "Change website domain",
	Long:              "Changes the primary domain for an Agency Plan website.\n\nProvide the current domain in the path and the new domain in the request body.\nSet domain to null to revert to the temporary domain.",
	Example:           "  hostinger agency-hosting domains change-website zpwlGlp19 old.example.com --domain new.example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := changeWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingChangeWebsiteDomainV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ChangeWebsiteCmd, utils.StringArg("website_uid"), utils.StringArg("from_domain"))
	ChangeWebsiteCmd.Flags().StringP("domain", "", "", "New domain to assign to the website. Set to null to revert to the temporary domain.")
	ChangeWebsiteCmd.MarkFlagRequired("domain")
	ChangeWebsiteCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func changeWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
