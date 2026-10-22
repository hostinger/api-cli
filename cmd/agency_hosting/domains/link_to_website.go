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

var LinkToWebsiteCmd = &cobra.Command{
	Use:               "link-to-website <website_uid>",
	Short:             "Link domain to website",
	Long:              "Links a domain to the specified Agency Plan website so it can serve traffic for that domain.",
	Example:           "  hostinger agency-hosting domains link-to-website zpwlGlp19 --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := linkToWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingLinkDomainToWebsiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(LinkToWebsiteCmd, utils.StringArg("website_uid"))
	LinkToWebsiteCmd.Flags().StringP("domain", "", "", "Fully qualified domain name to link to the website")
	LinkToWebsiteCmd.MarkFlagRequired("domain")
	LinkToWebsiteCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func linkToWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
