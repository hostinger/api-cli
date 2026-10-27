package site

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ErrNoUsername error = &utils.FlagError{Err: errors.New("no hosting account given: pass the username, or run `hostinger link`")}

// ResolveUsername: arg, else --username, HOSTINGER_USERNAME, the account of
// the default website, then a prompt.
func ResolveUsername(cmd *cobra.Command, args []string, extras int) (username string, rest []string, err error) {
	if len(args)-extras >= 1 {
		return args[0], args[1:], nil
	}
	rest = args
	if cmd != nil {
		if f := contextFlag(cmd, "username"); f != nil && f.Changed && f.Value.String() != "" {
			return f.Value.String(), rest, nil
		}
	}
	if v := os.Getenv("HOSTINGER_USERNAME"); v != "" {
		return v, rest, nil
	}
	ctx := context.Background()
	if cmd != nil && cmd.Context() != nil {
		ctx = cmd.Context()
	}
	domain, err := DefaultDomain(cmd)
	if err != nil {
		return "", nil, err
	}
	if p, _, _ := Find("."); p != nil && p.Username != "" && strings.EqualFold(p.Domain, domain) {
		return p.Username, rest, nil
	}
	if domain != "" {
		username, err = Username(ctx, domain)
		if err != nil {
			return "", nil, err
		}
		return username, rest, nil
	}
	if !Interactive() {
		return "", nil, ErrNoUsername
	}
	w, err := ChooseWebsite(ctx, "")
	if err != nil {
		return "", nil, err
	}
	return w.Username, rest, nil
}
