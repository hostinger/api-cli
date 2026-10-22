package horizons

import (
	"github.com/hostinger/api-cli/cmd/horizons/websites"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "horizons",
	Short: "Horizons commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(websites.GroupCmd)
}
