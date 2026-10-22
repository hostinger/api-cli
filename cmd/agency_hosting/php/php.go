package php

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "php",
	Short: "PHP commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ListExtensionsForWebsiteCmd)
	GroupCmd.AddCommand(ListOptionsForWebsiteCmd)
	GroupCmd.AddCommand(ListVersionsForOrderCmd)
	GroupCmd.AddCommand(ListVersionsForWebsiteCmd)
	GroupCmd.AddCommand(ReplaceWebsiteExtensionsCmd)
	GroupCmd.AddCommand(ReplaceWebsiteOptionsCmd)
	GroupCmd.AddCommand(UpdateWebsiteVersionCmd)
}
