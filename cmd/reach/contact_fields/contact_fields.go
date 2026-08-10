package contact_fields

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "contact-fields",
	Short: "Contact Fields commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(UpdateCmd)
}
