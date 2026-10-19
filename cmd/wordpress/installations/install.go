package installations

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

var InstallCmd = &cobra.Command{
	Use:   "install <username>",
	Short: "Install WordPress",
	Long:  "Install WordPress on an existing website.\n\nThe website must already exist before calling this endpoint. To create a new\nwebsite first, use POST /api/hosting/v1/orders/{order_id}/onboardings\n(`hosting_websites_start-setup`) and poll GET /api/hosting/v1/onboardings\n(`hosting_websites_list-setups`) with the `domain` filter until its `status` is\n`completed`.\n\nSend `type: headless` to turn a headless website (for example one created by\n`hosting_websites_start-setup` with `type: headless_empty`) into a headless WordPress\nwebsite: WordPress is installed into the `cms` directory with the Hostinger headless\nMU plugin and the website is marked as a headless WordPress website. `directory` is\nnot allowed with it. The target website must already be headless: the headless\ninstall rewrites the root `.htaccess`, so on any other website the request is\nrejected with 422.\n\nCall GET /api/hosting/v1/wordpress/installations (`wordpress_installations_list`)\nfiltered by username and domain before proceeding to check whether WordPress is\nalready installed on the target domain/path. If WordPress already exists and\n`overwrite` is false (the default), the async job will fail.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that WordPress is ready. Installation typically\ntakes 1-2 minutes. Poll GET /api/hosting/v1/wordpress/installations\n(`wordpress_installations_list`) filtered by username and domain to track\nprogress. When the installation appears in that list, WordPress is ready.\n\nReturns 422 when the account already uses all the databases its plan allows,\nunless `database.name` is an existing database on the account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "auto-updates", []string{"all", "none", "minor"})
		utils.EnumCheck(cmd, "type", []string{"default", "headless"})
		payload, err := json.Marshal(installBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingInstallWordPressV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	InstallCmd.Flags().StringP("auto-updates", "", "", "WordPress core auto-update policy (one of: all, none, minor)")
	InstallCmd.Flags().StringP("credentials", "", "", "WordPress admin credentials (JSON)")
	InstallCmd.Flags().StringP("database", "", "", "Optional. If the named database already exists on the account, it is used for this WordPress install. Otherwise a new database is created with this name, or with a generated name when database is omitted or null. A new database gets a random database user and counts toward the plan's database limit. (JSON)")
	InstallCmd.Flags().StringP("directory", "", "", "Installation subdirectory relative to the document root; omit for the root. Not allowed with `type: headless`, which always installs into `cms`.")
	InstallCmd.Flags().StringP("domain", "", "", "Domain of the existing website where WordPress will be installed")
	InstallCmd.Flags().StringP("language", "", "", "WordPress locale. Defaults to en_US when omitted.")
	InstallCmd.Flags().BoolP("overwrite", "", false, "When false (default), does not replace an existing installation. If WordPress is already installed on the domain/path, the async install job fails unless true.")
	InstallCmd.Flags().StringP("site-title", "", "", "Title of the WordPress site")
	InstallCmd.Flags().StringP("type", "", "", "Installation kind; `default` when omitted. `headless` installs WordPress into the `cms` directory with the Hostinger headless MU plugin and marks the website as a headless WordPress website. Accepted only for a website that is already headless, e.g. one created by `POST /api/hosting/v1/orders/{order_id}/onboardings` (`hosting_websites_start-setup`) with `type: headless_empty`; any other website is rejected with 422. (one of: default, headless)")
	InstallCmd.Flags().StringP("version", "", "", "WordPress core version to install. If omitted, the latest core version compatible with the account vhost PHP version is selected.")
	InstallCmd.MarkFlagRequired("credentials")
	InstallCmd.MarkFlagRequired("domain")
	InstallCmd.MarkFlagRequired("site-title")
}

func installBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("auto-updates") {
		v, _ := cmd.Flags().GetString("auto-updates")
		body["auto_updates"] = v
	}
	credentialsVal, _ := cmd.Flags().GetString("credentials")
	body["credentials"] = utils.JSONValue(credentialsVal, "credentials")
	if cmd.Flags().Changed("database") {
		v, _ := cmd.Flags().GetString("database")
		body["database"] = utils.JSONValue(v, "database")
	}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	if cmd.Flags().Changed("language") {
		v, _ := cmd.Flags().GetString("language")
		body["language"] = v
	}
	if cmd.Flags().Changed("overwrite") {
		v, _ := cmd.Flags().GetBool("overwrite")
		body["overwrite"] = v
	}
	siteTitleVal, _ := cmd.Flags().GetString("site-title")
	body["site_title"] = siteTitleVal
	if cmd.Flags().Changed("type") {
		v, _ := cmd.Flags().GetString("type")
		body["type"] = v
	}
	if cmd.Flags().Changed("version") {
		v, _ := cmd.Flags().GetString("version")
		body["version"] = v
	}
	return body
}
