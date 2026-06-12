package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PhpmyadminLinkCmd = &cobra.Command{
	Use:   "phpmyadmin-link <username> <name>",
	Short: "Get phpMyAdmin link",
	Long:  "Returns a direct sign-on link to phpMyAdmin for the specified database.\n\nUse this when a visual database interface is needed for SQL queries, imports, exports, or table management.\nThe database name must be the full name returned by the list databases endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetPhpMyAdminLinkV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
