package wordpress

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
	Long:  "Install WordPress on an existing website.\n\nThe website must already exist before calling this endpoint. To create a new\nwebsite first, use POST /api/hosting/v1/websites and poll\nGET /api/hosting/v1/websites until it appears.\n\nCall GET /api/hosting/v1/wordpress/installations filtered by username and\ndomain before proceeding to check whether WordPress is already installed on\nthe target domain/path. If WordPress already exists and `overwrite` is false\n(the default), the async job will fail.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that WordPress is ready. Installation typically\ntakes 1-2 minutes. Poll GET /api/hosting/v1/wordpress/installations filtered\nby username and domain to track progress. When the installation appears in\nthat list, WordPress is ready.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "auto-updates", []string{"all", "none", "minor"})
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
	InstallCmd.Flags().StringP("database", "", "", "Optional. If the named database already exists, it will be used for this WordPress install. Otherwise a new database is created with a generated name and random credentials. (JSON)")
	InstallCmd.Flags().StringP("directory", "", "", "Relative directory to install WordPress into. Defaults to the website root when omitted.")
	InstallCmd.Flags().StringP("domain", "", "", "Domain of the existing website where WordPress will be installed")
	InstallCmd.Flags().StringP("language", "", "", "WordPress locale. Defaults to en_US when omitted.")
	InstallCmd.Flags().BoolP("overwrite", "", false, "When false (default), does not replace an existing installation. If WordPress is already installed on the domain/path, the async install job fails unless true.")
	InstallCmd.Flags().StringP("site-title", "", "", "Title of the WordPress site")
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
	if cmd.Flags().Changed("version") {
		v, _ := cmd.Flags().GetString("version")
		body["version"] = v
	}
	return body
}
