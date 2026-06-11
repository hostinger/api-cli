package docker

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker project management",
	Long:  `Manage Docker Compose projects on a virtual machine. You can list projects, inspect their contents, containers and logs, and create, update, delete, start, stop or restart them.`,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ContainersCmd)
	GroupCmd.AddCommand(LogsCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(StartCmd)
	GroupCmd.AddCommand(StopCmd)
	GroupCmd.AddCommand(RestartCmd)
}
