package horizons

import (
	"github.com/hostinger/api-cli/cmd/horizons/websites"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "horizons",
	Short: "Horizons commands",
}

func init() {
	GroupCmd.AddCommand(websites.GroupCmd)
}
