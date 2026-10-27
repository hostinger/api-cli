package website_setups

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create [order-id]",
	Short:             "Create a new website",
	Long:              "Provisions a new website on one of your Agency Plan hosting orders.\n\nChoose the datacenter, stack (`flavor`), and PHP version for the site. Optionally attach\nyour own `domain` — omit it, set it to `null`, or leave it unavailable and a free\n`*.hostingersite.com` subdomain is generated instead — and/or install WordPress by\nsupplying the `wordpress` details (admin account, site title, and language).\n\nCommon setups:\n- Plain PHP site: `flavor` set to `php-fpm`, with `settings.php.version`; omit\n  `wordpress` and `type`.\n- WordPress site: `flavor` set to the desired WordPress version (e.g. `wp-7.0`), plus\n  the `wordpress` block (admin account, title, language).\n- Static/Node.js frontend app: `flavor` set to `php-fpm` and `type` set to\n  `node-static`.\n\nProvisioning runs in the background, so the response returns immediately with a setup UUID\nthat identifies the job. The new website becomes reachable once provisioning finishes.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --settings\n    {\n      \"php\": {\n        \"version\": \"8.3\"\n      }\n    }\n\n  --wordpress\n    {\n      \"admin\": {\n        \"email\": \"admin@example.com\",\n        \"password\": \"S3cureP@ssw0rd\",\n        \"user\": \"admin\"\n      },\n      \"language\": \"en_US\",\n      \"title\": \"My Blog\"\n    }",
	Example:           "  hostinger agency-hosting website-setups create 123456 --datacenter-code us-east --flavor wp-6.8.2 --settings '{\"php\":{\"version\":\"8.3\"}}'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "agency-hosting")
		if err != nil {
			return api.CommandError(err)
		}
		orderIDInt, err := utils.ScopeInt("order ID", orderID)
		if err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "type", []string{"node-static"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingCreateANewWebsiteV1WithBodyWithResponse(ctx, orderIDInt, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CreateCmd, 1, utils.IntArg("order_id"))
	CreateCmd.Flags().StringP("datacenter-code", "", "", "Datacenter code where the website should be provisioned. Available codes depend on live capacity and are not a fixed set.")
	CreateCmd.Flags().StringP("domain", "", "", "Primary domain to attach to the website. Omit or set to null to get a free auto-generated *.hostingersite.com subdomain instead.")
	CreateCmd.Flags().StringP("flavor", "", "", "Setup flavor: a specific WordPress version in the format `wp-<major>.<minor>` or `wp-<major>.<minor>.<patch>` (e.g. `wp-6.8.2`), or `php-fpm` for a plain PHP stack. Generic versions like `wp-latest` are not allowed.")
	CreateCmd.Flags().StringP("settings", "", "", "Website settings (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("type", "", "", "Website type (one of: node-static)")
	CreateCmd.Flags().StringP("wordpress", "", "", "WordPress installation options (JSON, @file or -; shape in the help above)")
	CreateCmd.MarkFlagRequired("datacenter-code")
	CreateCmd.MarkFlagRequired("flavor")
	CreateCmd.MarkFlagRequired("settings")
	CreateCmd.RegisterFlagCompletionFunc("datacenter-code", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("flavor", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("settings", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"node-static"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("wordpress", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	datacenterCodeVal, _ := cmd.Flags().GetString("datacenter-code")
	body["datacenter_code"] = datacenterCodeVal
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		body["domain"] = v
	}
	flavorVal, _ := cmd.Flags().GetString("flavor")
	body["flavor"] = flavorVal
	settingsVal, _ := cmd.Flags().GetString("settings")
	settingsValJSON, err := utils.JSONValue(settingsVal, "settings")
	if err != nil {
		return nil, err
	}
	body["settings"] = settingsValJSON
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
