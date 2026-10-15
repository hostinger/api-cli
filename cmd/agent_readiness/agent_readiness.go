package agent_readiness

import (
	"github.com/hostinger/api-cli/cmd/agent_readiness/checks"
	"github.com/hostinger/api-cli/cmd/agent_readiness/scans"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "agent-readiness",
	Short: "Agent Readiness commands",
}

func init() {
	GroupCmd.AddCommand(checks.GroupCmd)
	GroupCmd.AddCommand(scans.GroupCmd)
}
