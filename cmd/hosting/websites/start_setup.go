package websites

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

var StartSetupCmd = &cobra.Command{
	Use:   "start-setup <order_id>",
	Short: "Start website setup",
	Long:  "Starts a website setup on a Web or Cloud hosting order and returns the created setup\nright away; the website itself is provisioned asynchronously. Poll the list website\nsetups endpoint with the `domain` filter every 10 to 15 seconds and wait for\n`status: completed` before uploading files, deploying or creating databases.\n\nOmit `type` for an empty website. `type: wordpress` installs WordPress in the website\nroot with the admin user, email and password from `wordpress`, the domain as the site\ntitle, and `en_US` when `wordpress.language` is omitted. The headless types\n(`headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless\nwebsite; `headless_wordpress` additionally installs WordPress into the `cms` directory\nof the website root with generated credentials.\n\nOmit `domain` to set the website up on a generated temporary free subdomain.\n\nThe order must already have a hosting account: to create the first website on a new\nhosting plan use the create website endpoint, which takes the `datacenter_code`.\nReturns 404 when the order does not exist or is not accessible to the authenticated\nclient, and 409 with a `Retry-After` header while a setup for the same domain is still\nrunning.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "type", []string{"wordpress", "headless_wordpress", "headless_ecommerce", "headless_pocketbase"})
		payload, err := json.Marshal(startSetupBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingStartWebsiteSetupV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	StartSetupCmd.Flags().StringP("domain", "", "", "Customer-owned domain. Cannot start with \"www.\". Omit or `null` to set the website up on a generated temporary free subdomain.")
	StartSetupCmd.Flags().StringP("type", "", "", "Website type. Omit or `null` for an empty website. `wordpress` installs WordPress in the website root and requires `wordpress`. The headless types (`headless_wordpress`, `headless_ecommerce`, `headless_pocketbase`) create a headless website; `headless_wordpress` additionally installs WordPress into the `cms` directory with generated credentials. (one of: wordpress, headless_wordpress, headless_ecommerce, headless_pocketbase)")
	StartSetupCmd.Flags().StringP("wordpress", "", "", "WordPress install settings. Required when `type` is `wordpress`, not allowed otherwise. The site title is the domain. (JSON)")
}

func startSetupBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		body["domain"] = v
	}
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
