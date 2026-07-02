package themes

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
	Use:   "install <username> <software>",
	Short: "Install WordPress theme",
	Long:  "Install a theme on an existing WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield).\n\nWhen the theme is one of the Hostinger themes (hostinger-blog,\nhostinger-affiliate-theme, hostinger-ai-theme), the optional `palette`,\n`layout`, and `font` fields are forwarded to the custom installer (defaults:\npalette1, layout1, default). For any other theme they are ignored.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that the theme is ready.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "font", []string{"professional", "modern", "elegant", "creative", "dynamic", "default"})
		payload, err := json.Marshal(installBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingInstallWordPressThemeV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	InstallCmd.Flags().StringP("font", "", "default", "Font identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted. (one of: professional, modern, elegant, creative, dynamic, default)")
	InstallCmd.Flags().StringP("layout", "", "layout1", "Layout identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted.")
	InstallCmd.Flags().StringP("palette", "", "palette1", "Palette identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted.")
	InstallCmd.Flags().StringP("theme", "", "", "Slug of the theme to install. Hostinger theme slugs (hostinger-blog, hostinger-affiliate-theme, hostinger-ai-theme) trigger the custom installer and forward the optional palette/layout/font fields; any other WordPress theme slug uses the standard installer and ignores those fields.")
	InstallCmd.MarkFlagRequired("theme")
}

func installBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("font") {
		v, _ := cmd.Flags().GetString("font")
		body["font"] = v
	}
	if cmd.Flags().Changed("layout") {
		v, _ := cmd.Flags().GetString("layout")
		body["layout"] = v
	}
	if cmd.Flags().Changed("palette") {
		v, _ := cmd.Flags().GetString("palette")
		body["palette"] = v
	}
	themeVal, _ := cmd.Flags().GetString("theme")
	body["theme"] = themeVal
	return body
}
