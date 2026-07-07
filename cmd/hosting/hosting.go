package hosting

import (
	"github.com/hostinger/api-cli/cmd/hosting/cache"
	"github.com/hostinger/api-cli/cmd/hosting/cron_jobs"
	"github.com/hostinger/api-cli/cmd/hosting/databases"
	"github.com/hostinger/api-cli/cmd/hosting/datacenters"
	"github.com/hostinger/api-cli/cmd/hosting/domains"
	"github.com/hostinger/api-cli/cmd/hosting/nodejs"
	"github.com/hostinger/api-cli/cmd/hosting/orders"
	"github.com/hostinger/api-cli/cmd/hosting/php"
	"github.com/hostinger/api-cli/cmd/hosting/websites"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "hosting",
	Short: "Hosting commands",
}

func init() {
	GroupCmd.AddCommand(cache.GroupCmd)
	GroupCmd.AddCommand(cron_jobs.GroupCmd)
	GroupCmd.AddCommand(databases.GroupCmd)
	GroupCmd.AddCommand(datacenters.GroupCmd)
	GroupCmd.AddCommand(domains.GroupCmd)
	GroupCmd.AddCommand(nodejs.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
	GroupCmd.AddCommand(php.GroupCmd)
	GroupCmd.AddCommand(websites.GroupCmd)
}
