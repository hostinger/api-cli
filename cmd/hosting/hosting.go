package hosting

import (
	"github.com/hostinger/api-cli/cmd/hosting/databases"
	"github.com/hostinger/api-cli/cmd/hosting/datacenters"
	"github.com/hostinger/api-cli/cmd/hosting/domains"
	"github.com/hostinger/api-cli/cmd/hosting/nodejs"
	"github.com/hostinger/api-cli/cmd/hosting/orders"
	"github.com/hostinger/api-cli/cmd/hosting/websites"
	"github.com/hostinger/api-cli/cmd/hosting/wordpress"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "hosting",
	Short: "Hosting commands",
}

func init() {
	GroupCmd.AddCommand(databases.GroupCmd)
	GroupCmd.AddCommand(datacenters.GroupCmd)
	GroupCmd.AddCommand(domains.GroupCmd)
	GroupCmd.AddCommand(nodejs.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
	GroupCmd.AddCommand(websites.GroupCmd)
	GroupCmd.AddCommand(wordpress.GroupCmd)
}
