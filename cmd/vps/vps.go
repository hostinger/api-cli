package vps

import (
	"github.com/hostinger/api-cli/cmd/vps/actions"
	"github.com/hostinger/api-cli/cmd/vps/backups"
	"github.com/hostinger/api-cli/cmd/vps/data_centers"
	"github.com/hostinger/api-cli/cmd/vps/docker"
	"github.com/hostinger/api-cli/cmd/vps/firewall"
	"github.com/hostinger/api-cli/cmd/vps/monarx"
	"github.com/hostinger/api-cli/cmd/vps/post_install_scripts"
	"github.com/hostinger/api-cli/cmd/vps/ptr"
	"github.com/hostinger/api-cli/cmd/vps/public_keys"
	"github.com/hostinger/api-cli/cmd/vps/recovery"
	"github.com/hostinger/api-cli/cmd/vps/snapshots"
	"github.com/hostinger/api-cli/cmd/vps/templates"
	"github.com/hostinger/api-cli/cmd/vps/virtual_machines"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "vps",
	Short: "VPS commands",
}

func init() {
	GroupCmd.AddCommand(actions.GroupCmd)
	GroupCmd.AddCommand(backups.GroupCmd)
	GroupCmd.AddCommand(data_centers.GroupCmd)
	GroupCmd.AddCommand(docker.GroupCmd)
	GroupCmd.AddCommand(firewall.GroupCmd)
	GroupCmd.AddCommand(monarx.GroupCmd)
	GroupCmd.AddCommand(post_install_scripts.GroupCmd)
	GroupCmd.AddCommand(ptr.GroupCmd)
	GroupCmd.AddCommand(public_keys.GroupCmd)
	GroupCmd.AddCommand(recovery.GroupCmd)
	GroupCmd.AddCommand(snapshots.GroupCmd)
	GroupCmd.AddCommand(templates.GroupCmd)
	GroupCmd.AddCommand(virtual_machines.GroupCmd)
}
