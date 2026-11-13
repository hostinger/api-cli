package gpu

import (
	"github.com/hostinger/api-cli/cmd/gpu/catalog"
	"github.com/hostinger/api-cli/cmd/gpu/exposed_services"
	"github.com/hostinger/api-cli/cmd/gpu/instances"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "gpu",
	Short: "GPU commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(catalog.GroupCmd)
	GroupCmd.AddCommand(exposed_services.GroupCmd)
	GroupCmd.AddCommand(instances.GroupCmd)
}
