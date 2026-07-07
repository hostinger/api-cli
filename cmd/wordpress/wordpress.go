package wordpress

import (
	"github.com/hostinger/api-cli/cmd/wordpress/ai_tools"
	"github.com/hostinger/api-cli/cmd/wordpress/installations"
	"github.com/hostinger/api-cli/cmd/wordpress/litespeed_cache"
	"github.com/hostinger/api-cli/cmd/wordpress/login"
	"github.com/hostinger/api-cli/cmd/wordpress/maintenance"
	"github.com/hostinger/api-cli/cmd/wordpress/object_cache"
	"github.com/hostinger/api-cli/cmd/wordpress/plugins"
	"github.com/hostinger/api-cli/cmd/wordpress/themes"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "WordPress commands",
}

func init() {
	GroupCmd.AddCommand(ai_tools.GroupCmd)
	GroupCmd.AddCommand(installations.GroupCmd)
	GroupCmd.AddCommand(litespeed_cache.GroupCmd)
	GroupCmd.AddCommand(login.GroupCmd)
	GroupCmd.AddCommand(maintenance.GroupCmd)
	GroupCmd.AddCommand(object_cache.GroupCmd)
	GroupCmd.AddCommand(plugins.GroupCmd)
	GroupCmd.AddCommand(themes.GroupCmd)
}
