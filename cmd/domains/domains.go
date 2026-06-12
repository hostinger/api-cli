package domains

import (
	"github.com/hostinger/api-cli/cmd/domains/availability"
	"github.com/hostinger/api-cli/cmd/domains/forwarding"
	"github.com/hostinger/api-cli/cmd/domains/portfolio"
	"github.com/hostinger/api-cli/cmd/domains/verifications"
	"github.com/hostinger/api-cli/cmd/domains/whois"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "domains",
	Short: "Domains commands",
}

func init() {
	GroupCmd.AddCommand(availability.GroupCmd)
	GroupCmd.AddCommand(forwarding.GroupCmd)
	GroupCmd.AddCommand(portfolio.GroupCmd)
	GroupCmd.AddCommand(verifications.GroupCmd)
	GroupCmd.AddCommand(whois.GroupCmd)
}
