package post_install_scripts

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "post-install-scripts",
	Short: "Post-install scripts commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
