package login

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "login",
	Short: "Login commands",
}

func init() {
	GroupCmd.AddCommand(CreateLinksCmd)
}
