package agency_hosting

import (
	"github.com/hostinger/api-cli/cmd/agency_hosting/cache"
	"github.com/hostinger/api-cli/cmd/agency_hosting/cron_jobs"
	"github.com/hostinger/api-cli/cmd/agency_hosting/datacenters"
	"github.com/hostinger/api-cli/cmd/agency_hosting/domains"
	"github.com/hostinger/api-cli/cmd/agency_hosting/files"
	"github.com/hostinger/api-cli/cmd/agency_hosting/website_setups"
	"github.com/hostinger/api-cli/cmd/agency_hosting/websites"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "agency-hosting",
	Short: "Agency Hosting commands",
}

func init() {
	GroupCmd.AddCommand(cache.GroupCmd)
	GroupCmd.AddCommand(cron_jobs.GroupCmd)
	GroupCmd.AddCommand(datacenters.GroupCmd)
	GroupCmd.AddCommand(domains.GroupCmd)
	GroupCmd.AddCommand(files.GroupCmd)
	GroupCmd.AddCommand(website_setups.GroupCmd)
	GroupCmd.AddCommand(websites.GroupCmd)
}
