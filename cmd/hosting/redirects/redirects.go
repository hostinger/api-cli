package redirects

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "redirects",
	Short: "Redirects commands",
}

func init() {
	GroupCmd.AddCommand(CreateWebsiteCmd)
	GroupCmd.AddCommand(DeleteWebsiteCmd)
	GroupCmd.AddCommand(ListWebsiteCmd)
}
