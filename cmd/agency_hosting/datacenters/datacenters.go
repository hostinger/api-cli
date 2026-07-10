package datacenters

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "datacenters",
	Short: "Datacenters commands",
}

func init() {
	GroupCmd.AddCommand(ListForPlanOrderCmd)
}
