package website_setups

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

var CreateCmd = &cobra.Command{
	Use:   "create <order_id>",
	Short: "Create a new website",
	Long:  "Provisions a new website on one of your Agency Plan hosting orders.\n\nChoose the datacenter, stack (`flavor`), and PHP version for the site. Optionally attach\nyour own `domain` — omit it, set it to `null`, or leave it unavailable and a free\n`*.hostingersite.com` subdomain is generated instead — and/or install WordPress by\nsupplying the `wordpress` details (admin account, site title, and language).\n\nCommon setups:\n- **Plain PHP site**: `flavor` set to `php-fpm`, with `settings.php.version`; omit\n  `wordpress` and `type`.\n- **WordPress site**: `flavor` set to the desired WordPress version (e.g. `wp-7.0`), plus\n  the `wordpress` block (admin account, title, language).\n- **Static/Node.js frontend app**: `flavor` set to `php-fpm` and `type` set to\n  `node-static`.\n\nProvisioning runs in the background, so the response returns immediately with a setup UUID\nthat identifies the job. The new website becomes reachable once provisioning finishes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"horizons", "node-static"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingCreateANewWebsiteV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("clone", "", "", "Clone the new website from an existing website (JSON)")
	CreateCmd.Flags().StringP("datacenter-code", "", "", "Datacenter code where the website should be provisioned. Available codes depend on live capacity and are not a fixed set.")
	CreateCmd.Flags().StringP("derive-domain", "", "", "Derive the domain from an existing vhost (JSON)")
	CreateCmd.Flags().StringP("domain", "", "", "Primary domain to attach to the website. Omit or set to null to get a free auto-generated *.hostingersite.com subdomain instead.")
	CreateCmd.Flags().StringP("flavor", "", "", "Setup flavor: a specific WordPress version in the format `wp-<major>.<minor>` or `wp-<major>.<minor>.<patch>` (e.g. `wp-6.8.2`), or `php-fpm` for a plain PHP stack. Generic versions like `wp-latest` are not allowed.")
	CreateCmd.Flags().StringP("settings", "", "", "Website settings (JSON)")
	CreateCmd.Flags().StringP("type", "", "", "Website type (one of: horizons, node-static)")
	CreateCmd.Flags().StringP("wordpress", "", "", "WordPress installation options (JSON)")
	CreateCmd.MarkFlagRequired("datacenter-code")
	CreateCmd.MarkFlagRequired("flavor")
	CreateCmd.MarkFlagRequired("settings")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("clone") {
		v, _ := cmd.Flags().GetString("clone")
		body["clone"] = utils.JSONValue(v, "clone")
	}
	datacenterCodeVal, _ := cmd.Flags().GetString("datacenter-code")
	body["datacenter_code"] = datacenterCodeVal
	if cmd.Flags().Changed("derive-domain") {
		v, _ := cmd.Flags().GetString("derive-domain")
		body["derive_domain"] = utils.JSONValue(v, "derive-domain")
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		body["domain"] = v
	}
	flavorVal, _ := cmd.Flags().GetString("flavor")
	body["flavor"] = flavorVal
	settingsVal, _ := cmd.Flags().GetString("settings")
	body["settings"] = utils.JSONValue(settingsVal, "settings")
	if cmd.Flags().Changed("type") {
		v, _ := cmd.Flags().GetString("type")
		body["type"] = v
	}
	if cmd.Flags().Changed("wordpress") {
		v, _ := cmd.Flags().GetString("wordpress")
		body["wordpress"] = utils.JSONValue(v, "wordpress")
	}
	return body
}
