package websites

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "websites",
	Short: "Websites commands",
}

func init() {
	GroupCmd.AddCommand(BuildNodejsAssetsCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListProcessesCmd)
}
