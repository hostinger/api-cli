package php

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "php",
	Short: "PHP commands",
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
