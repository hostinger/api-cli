package availability

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "availability",
	Short: "Availability commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CheckCmd)
	GroupCmd.AddCommand(SuggestNamesFromCmd)
	GroupCmd.AddCommand(SuggestNamesFromDescriptionCmd)
}
