package ecommerce

import (
	"github.com/hostinger/api-cli/cmd/ecommerce/stores"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "ecommerce",
	Short: "Ecommerce commands",
}

func init() {
	GroupCmd.AddCommand(stores.GroupCmd)
}
