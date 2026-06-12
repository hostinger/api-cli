package docker

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker Manager commands",
}

func init() {
	GroupCmd.AddCommand(ContainersCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(LogsCmd)
	GroupCmd.AddCommand(RestartCmd)
	GroupCmd.AddCommand(StartCmd)
	GroupCmd.AddCommand(StopCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
