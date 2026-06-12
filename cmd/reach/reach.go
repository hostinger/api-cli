package reach

import (
	"github.com/hostinger/api-cli/cmd/reach/contacts"
	"github.com/hostinger/api-cli/cmd/reach/profiles"
	"github.com/hostinger/api-cli/cmd/reach/segments"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "reach",
	Short: "Reach commands",
}

func init() {
	GroupCmd.AddCommand(contacts.GroupCmd)
	GroupCmd.AddCommand(profiles.GroupCmd)
	GroupCmd.AddCommand(segments.GroupCmd)
}
