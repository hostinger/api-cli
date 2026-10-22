package utils

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	cobra.AddTemplateFunc("requiredFlagUsages", func(c *cobra.Command) string { return flagUsages(c, true) })
	cobra.AddTemplateFunc("optionalFlagUsages", func(c *cobra.Command) string { return flagUsages(c, false) })
}

// UsageTemplate is cobra's default usage template with the local flags split
// into "Required Flags" and "Flags".
const UsageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}{{$required := requiredFlagUsages .}}{{if $required}}

Required Flags:
{{$required}}{{end}}{{$optional := optionalFlagUsages .}}{{if $optional}}

Flags:
{{$optional}}{{end}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// flagUsages renders the command's local flags that are (or are not) required.
func flagUsages(c *cobra.Command, required bool) string {
	fs := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
	fs.SortFlags = c.LocalFlags().SortFlags
	c.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Hidden && IsRequired(f) == required {
			fs.AddFlag(f)
		}
	})
	return strings.TrimRight(fs.FlagUsages(), " \t\n")
}
