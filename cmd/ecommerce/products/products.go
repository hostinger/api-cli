package products

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "products",
	Short: "Products commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateDigitalCmd)
	GroupCmd.AddCommand(CreateImageUploadUrlCmd)
	GroupCmd.AddCommand(CreatePhysicalCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
	GroupCmd.AddCommand(UploadAndAttachImageCmd)
}
