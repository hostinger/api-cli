package segments

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "segments",
	Short: "Segments commands",
}

func init() {
	GroupCmd.AddCommand(CountProfileContactsCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(CreateProfileCmd)
	GroupCmd.AddCommand(DeleteProfileCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListContactsCmd)
	GroupCmd.AddCommand(ListProfileCmd)
	GroupCmd.AddCommand(ListProfileContactsCmd)
	GroupCmd.AddCommand(ProfileCmd)
	GroupCmd.AddCommand(UpdateProfileCmd)
}
