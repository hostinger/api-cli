package utils

import (
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// ArgKind is the type a positional argument must parse as.
type ArgKind int

const (
	ArgString ArgKind = iota
	ArgInt
	ArgUUID
)

// ArgSpec describes one positional argument of a generated command.
type ArgSpec struct {
	Name string // e.g. "virtual-machine-id"
	Kind ArgKind
}

func StringArg(name string) ArgSpec { return ArgSpec{Name: name, Kind: ArgString} }
func IntArg(name string) ArgSpec    { return ArgSpec{Name: name, Kind: ArgInt} }
func UUIDArg(name string) ArgSpec   { return ArgSpec{Name: name, Kind: ArgUUID} }

var (
	argSpecsMu sync.Mutex
	argSpecs   = map[*cobra.Command][]ArgSpec{}
	// argScopes: leading args a command resolves from context when left out.
	argScopes = map[*cobra.Command]int{}
	// promptedArgs holds the full argument list after interactive prompting,
	// which runs in the root's PersistentPreRunE where the slice cobra passes
	// to RunE can no longer grow.
	promptedArgs = map[*cobra.Command][]string{}
)

// RegisterArgs installs the positional-argument validator for cmd: exactly
// len(specs) arguments, each parsing as its kind. Missing trailing arguments
// are allowed in an interactive session; they are prompted for before RunE.
func RegisterArgs(cmd *cobra.Command, specs ...ArgSpec) {
	registerArgs(cmd, 0, specs)
}

// RegisterScopedArgs is RegisterArgs where the first scope args may be left
// out (the resolver supplies them); given args fill the specs from the end.
func RegisterScopedArgs(cmd *cobra.Command, scope int, specs ...ArgSpec) {
	registerArgs(cmd, scope, specs)
}

func registerArgs(cmd *cobra.Command, scope int, specs []ArgSpec) {
	argSpecsMu.Lock()
	argSpecs[cmd] = specs
	argScopes[cmd] = scope
	argSpecsMu.Unlock()
	cmd.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) > len(specs) {
			if len(specs) == 0 {
				return FlagErrorf("%q takes no arguments, got %q", cmd.CommandPath(), args)
			}
			if scope > 0 {
				return FlagErrorf("%q takes at most %d argument(s) (%s), got %d", cmd.CommandPath(), len(specs), argNames(specs), len(args))
			}
			return FlagErrorf("%q takes %d argument(s) (%s), got %d", cmd.CommandPath(), len(specs), argNames(specs), len(args))
		}
		if len(args) < len(specs)-scope && !Interactive(cmd) {
			return FlagErrorf("missing argument <%s> (usage: %s)", specs[scope+len(args)].Name, cmd.UseLine())
		}
		for i, a := range args {
			if err := checkArg(specs[argOffset(specs, scope, args)+i], a); err != nil {
				return err
			}
		}
		return nil
	}
}

// argOffset is the spec index of args[0].
func argOffset(specs []ArgSpec, scope int, args []string) int {
	return min(scope, len(specs)-len(args))
}

// ArgSpecs returns the positional arguments registered for cmd.
func ArgSpecs(cmd *cobra.Command) []ArgSpec {
	argSpecsMu.Lock()
	defer argSpecsMu.Unlock()
	return argSpecs[cmd]
}

func ArgScope(cmd *cobra.Command) int {
	argSpecsMu.Lock()
	defer argSpecsMu.Unlock()
	return argScopes[cmd]
}

// PromptedArgs returns args extended with any values prompted for interactively.
func PromptedArgs(cmd *cobra.Command, args []string) []string {
	argSpecsMu.Lock()
	defer argSpecsMu.Unlock()
	if full := promptedArgs[cmd]; full != nil {
		return full
	}
	return args
}

// setPromptedArgs records (or, with nil, clears) the prompted argument list.
func setPromptedArgs(cmd *cobra.Command, args []string) {
	argSpecsMu.Lock()
	if args == nil {
		delete(promptedArgs, cmd)
	} else {
		promptedArgs[cmd] = args
	}
	argSpecsMu.Unlock()
}

func checkArg(spec ArgSpec, v string) error {
	switch spec.Kind {
	case ArgInt:
		if _, err := strconv.Atoi(v); err != nil {
			return FlagErrorf("invalid <%s> %q: expected an integer", spec.Name, v)
		}
	case ArgUUID:
		if _, err := uuid.Parse(v); err != nil {
			return FlagErrorf("invalid <%s> %q: expected a UUID", spec.Name, v)
		}
	}
	return nil
}

func argNames(specs []ArgSpec) string {
	s := ""
	for i, a := range specs {
		if i > 0 {
			s += " "
		}
		s += "<" + a.Name + ">"
	}
	return s
}

// ParseIntArg parses an integer argument where a bad value must not panic (shell
// completion of later arguments).
func ParseIntArg(s string) (int, error) { return strconv.Atoi(s) }

// ParseUUIDArg parses a UUID argument where a bad value must not panic.
func ParseUUIDArg(s string) (uuid.UUID, error) { return uuid.Parse(s) }
