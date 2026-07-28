package api_tokens

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "api-tokens",
	Short: "API Tokens commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RevokeCmd)
}
