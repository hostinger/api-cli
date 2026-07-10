package website_setups

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "website-setups",
	Short: "Website Setups commands",
}

func init() {
	GroupCmd.AddCommand(PlanStatusCmd)
	GroupCmd.AddCommand(ProvisionPlanCmd)
}
