package websites

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StartSetupCmd = &cobra.Command{
	Use:               "start-setup <order_id>",
	Short:             "Start website setup",
	Long:              "Starts a website setup on a Web or Cloud hosting order and returns the created setup\nright away; the website itself is provisioned asynchronously. Poll the list website\nsetups endpoint (`hosting_websites_list-setups`) with the `domain` filter every 10 to 15\nseconds and wait for `status: completed` before uploading files, deploying or creating\ndatabases.\n\nOmit `type` for an empty website. `type: wordpress` installs WordPress in the website\nroot with the admin user, email and password from `wordpress`, the domain as the site\ntitle, and `en_US` when `wordpress.language` is omitted. The headless types\n(`headless_empty`, `headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`)\ncreate a headless website; `headless_wordpress` additionally installs WordPress into the\n`cms` directory of the website root with generated credentials. `headless_empty` installs\nnothing: WordPress can be added later through\n`POST /api/hosting/v1/accounts/{username}/wordpress/installations`\n(`wordpress_installations_install`) with `type: headless`.\n\nOmit `domain` to set the website up on a generated temporary free subdomain. To know the\nfree subdomain before the call, so a timed-out call can be looked up and retried on the\nsame domain, generate one through `POST /api/hosting/v1/domains/free-subdomains`\n(`hosting_domains_generate-free-subdomain`) and send it as `domain`. It is accepted for\n24 hours, only from the client that generated it.\n\nWhen the order has no hosting account yet, the setup creates one first, in the datacenter\ngiven by `datacenter_code` or the order's default datacenter when it is omitted.\nReturns 404 when the order does not exist or is not accessible to the authenticated\nclient, and 409 with a `Retry-After` header while a setup for the same domain is still\nrunning.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --wordpress\n    {\n      \"admin\": {\n        \"email\": \"admin@example.com\",\n        \"password\": \"S3curePass123\",\n        \"user\": \"site_admin\"\n      },\n      \"is_ai_builder\": false,\n      \"language\": \"en_US\"\n    }",
	Example:           "  hostinger hosting websites start-setup 12345",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "type", []string{"wordpress", "headless_empty", "headless_wordpress", "headless_ecommerce", "headless_pocketbase"}); err != nil {
			return err
		}
		body, err := startSetupBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingStartWebsiteSetupV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(StartSetupCmd, utils.IntArg("order_id"))
	StartSetupCmd.Flags().StringP("datacenter-code", "", "", "Datacenter for the hosting account when the order does not have one yet. Omit to use the order's default datacenter. Ignored when the order already has a hosting account. List available datacenters to find valid codes.")
	StartSetupCmd.Flags().StringP("domain", "", "", "Customer-owned domain, or a free subdomain you generated through `POST /api/hosting/v1/domains/free-subdomains` (`hosting_domains_generate-free-subdomain`) in the last 24 hours. Cannot start with \"www.\" and can have at most one subdomain level, e.g. `blog.example.com`. Omit or `null` to set the website up on a generated temporary free subdomain.")
	StartSetupCmd.Flags().StringP("type", "", "", "Website type. Omit or `null` for an empty website. `wordpress` installs WordPress in the website root and requires `wordpress`. The headless types (`headless_empty`, `headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless website; `headless_wordpress` additionally installs WordPress into the `cms` directory with generated credentials. `headless_empty` installs nothing, so WordPress can be added later through `POST /api/hosting/v1/accounts/{username}/wordpress/installations` (`wordpress_installations_install`) with `type: headless`. (one of: wordpress, headless_empty, headless_wordpress, headless_ecommerce, headless_pocketbase)")
	StartSetupCmd.Flags().StringP("wordpress", "", "", "WordPress install settings. Required when `type` is `wordpress`, not allowed otherwise. The site title is the domain. (JSON, @file or -; shape in the help above)")
	StartSetupCmd.RegisterFlagCompletionFunc("datacenter-code", cobra.NoFileCompletions)
	StartSetupCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	StartSetupCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"wordpress", "headless_empty", "headless_wordpress", "headless_ecommerce", "headless_pocketbase"}, cobra.ShellCompDirectiveNoFileComp))
	StartSetupCmd.RegisterFlagCompletionFunc("wordpress", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func startSetupBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("datacenter-code") {
		v, _ := cmd.Flags().GetString("datacenter-code")
		body["datacenter_code"] = v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		body["domain"] = v
	}
	if cmd.Flags().Changed("type") {
		v, _ := cmd.Flags().GetString("type")
		body["type"] = v
	}
	if cmd.Flags().Changed("wordpress") {
		wordpressVal, _ := cmd.Flags().GetString("wordpress")
		wordpressValJSON, err := utils.JSONValue(wordpressVal, "wordpress")
		if err != nil {
			return nil, err
		}
		body["wordpress"] = wordpressValJSON
	}
	return body, nil
}
