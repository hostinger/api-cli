package cmd

import (
	"fmt"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.CLIVersion.String(true))
	},
}
