package themes

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:     "install <username> <software>",
	Short:   "Install WordPress theme",
	Long:    "Install a theme on an existing WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield).\n\nWhen the theme is one of the Hostinger themes (hostinger-blog,\nhostinger-affiliate-theme, hostinger-ai-theme), the optional `palette`,\n`layout`, and `font` fields are forwarded to the custom installer (defaults:\npalette1, layout1, default). For any other theme they are ignored.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that the theme is ready.",
	Example: "  hostinger wordpress themes install u123456789 1232456789 --theme hostinger-blog",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "font", []string{"professional", "modern", "elegant", "creative", "dynamic", "default"}); err != nil {
			return err
		}
		body, err := installBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingInstallWordPressThemeV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(InstallCmd, utils.StringArg("username"), utils.StringArg("software"))
	InstallCmd.Flags().StringP("font", "", "default", "Font identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted. (one of: professional, modern, elegant, creative, dynamic, default)")
	InstallCmd.Flags().StringP("layout", "", "layout1", "Layout identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted.")
	InstallCmd.Flags().StringP("palette", "", "palette1", "Palette identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted.")
	InstallCmd.Flags().StringP("theme", "", "", "Slug of the theme to install. Hostinger theme slugs (hostinger-blog, hostinger-affiliate-theme, hostinger-ai-theme) trigger the custom installer and forward the optional palette/layout/font fields; any other WordPress theme slug uses the standard installer and ignores those fields.")
	InstallCmd.MarkFlagRequired("theme")
	InstallCmd.RegisterFlagCompletionFunc("font", cobra.FixedCompletions([]cobra.Completion{"professional", "modern", "elegant", "creative", "dynamic", "default"}, cobra.ShellCompDirectiveNoFileComp))
	InstallCmd.RegisterFlagCompletionFunc("layout", cobra.NoFileCompletions)
	InstallCmd.RegisterFlagCompletionFunc("palette", cobra.NoFileCompletions)
	InstallCmd.RegisterFlagCompletionFunc("theme", cobra.NoFileCompletions)
}

func installBody(cmd *cobra.Command) (map[string]any, error) {
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
	return body, nil
}
