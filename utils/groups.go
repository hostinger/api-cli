package utils

import (
	"strings"

	"github.com/spf13/cobra"
)

// GroupRunE makes a command group fail on an unknown subcommand (with
// suggestions) instead of printing help and exiting 0. Without arguments it
// shows the group's help.
func GroupRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	if cmd.SuggestionsMinimumDistance <= 0 {
		cmd.SuggestionsMinimumDistance = 2 // cobra only defaults this on the root
	}
	msg := "unknown command " + `"` + args[0] + `"` + " for " + `"` + cmd.CommandPath() + `"`
	if s := cmd.SuggestionsFor(args[0]); len(s) > 0 {
		msg += "\n\nDid you mean this?\n\t" + strings.Join(s, "\n\t")
	}
	return FlagErrorf("%s", msg)
}
