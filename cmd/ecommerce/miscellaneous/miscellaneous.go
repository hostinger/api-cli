package miscellaneous

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "miscellaneous",
	Short: "Miscellaneous commands",
}

func init() {
	GroupCmd.AddCommand(CustomStorefrontSetupInstructionsCmd)
}
