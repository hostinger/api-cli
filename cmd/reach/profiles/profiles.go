package profiles

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Profiles commands",
}

func init() {
	GroupCmd.AddCommand(DomainDnsStatusCmd)
	GroupCmd.AddCommand(ListCmd)
}
