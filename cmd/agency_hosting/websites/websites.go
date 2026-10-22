package websites

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "websites",
	Short: "Websites commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(BuildNodejsAssetsCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListPlanCmd)
	GroupCmd.AddCommand(ListProcessesCmd)
}
