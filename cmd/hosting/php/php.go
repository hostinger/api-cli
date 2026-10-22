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
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(InfoCmd)
	GroupCmd.AddCommand(ResetExtensionsCmd)
	GroupCmd.AddCommand(UpdateExtensionsCmd)
	GroupCmd.AddCommand(UpdateOptionsCmd)
	GroupCmd.AddCommand(UpdateVersionCmd)
}
