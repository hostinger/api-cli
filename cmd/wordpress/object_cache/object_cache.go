package object_cache

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "object-cache",
	Short: "Object Cache commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ShowMemcachedStatusCmd)
	GroupCmd.AddCommand(ToggleMemcachedCmd)
}
