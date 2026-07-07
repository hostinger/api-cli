package object_cache

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "object-cache",
	Short: "Object Cache commands",
}

func init() {
	GroupCmd.AddCommand(ShowMemcachedStatusCmd)
	GroupCmd.AddCommand(ToggleMemcachedCmd)
}
