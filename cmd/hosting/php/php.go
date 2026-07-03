package php

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "php",
	Short: "PHP commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(InfoCmd)
	GroupCmd.AddCommand(ResetExtensionsCmd)
	GroupCmd.AddCommand(UpdateExtensionsCmd)
	GroupCmd.AddCommand(UpdateOptionsCmd)
	GroupCmd.AddCommand(UpdateVersionCmd)
}
