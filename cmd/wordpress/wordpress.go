package wordpress

import (
	"github.com/hostinger/api-cli/cmd/wordpress/installations"
	"github.com/hostinger/api-cli/cmd/wordpress/plugins"
	"github.com/hostinger/api-cli/cmd/wordpress/themes"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "WordPress commands",
}

func init() {
	GroupCmd.AddCommand(installations.GroupCmd)
	GroupCmd.AddCommand(plugins.GroupCmd)
	GroupCmd.AddCommand(themes.GroupCmd)
}
