package whois

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "whois",
	Short: "WHOIS commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UsageCmd)
}
