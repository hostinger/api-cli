package site

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ErrNoDomain error = &utils.FlagError{Err: errors.New("no domain given: pass the domain or --domain")}

// ResolveDomain: for hosting the domain is a website (see Resolve). For
// domains and DNS the link only counts when its domain is in the portfolio,
// since a free subdomain is not a zone.
func ResolveDomain(cmd *cobra.Command, args []string, extras int, group string) (domain string, rest []string, err error) {
	if len(args)-extras >= 1 {
		return NormalizeDomain(args[0]), args[1:], nil
	}
	rest = args
	if group == "hosting" {
		_, domain, rest, err = Resolve(cmd, args, extras)
		return domain, rest, err
	}
	if cmd != nil {
		if f := contextFlag(cmd, "domain"); f != nil && f.Changed && f.Value.String() != "" {
			return NormalizeDomain(f.Value.String()), rest, nil
		}
	}
	if v := os.Getenv("HOSTINGER_DOMAIN"); v != "" {
		return NormalizeDomain(v), rest, nil
	}
	ctx := context.Background()
	if cmd != nil && cmd.Context() != nil {
		ctx = cmd.Context()
	}
	p, _, err := Find(".")
	if err != nil {
		return "", nil, linkError(err)
	}
	// The portfolio is only needed to vet the link or to ask.
	var owned []string
	if p != nil || Interactive() {
		if owned, err = Portfolio(ctx); err != nil {
			return "", nil, err
		}
	}
	if p != nil {
		for _, d := range owned {
			if strings.EqualFold(d, p.Domain) {
				return d, rest, nil
			}
		}
	}
	if v := viper.GetString("default_domain"); v != "" {
		return NormalizeDomain(v), rest, nil
	}
	if !Interactive() {
		return "", nil, ErrNoDomain
	}
	switch len(owned) {
	case 0:
		return "", nil, errors.New("this account owns no domains")
	case 1:
		return owned[0], rest, nil
	}
	picked, err := pick(pterm.DefaultInteractiveSelect.WithOptions(owned).WithFilter(len(owned) > 8).WithMaxHeight(SelectHeight(len(owned))), "Which domain?")
	if err != nil {
		return "", nil, err
	}
	return picked, rest, nil
}

func Portfolio(ctx context.Context) ([]string, error) {
	var rows []struct {
		Domain string `json:"domain"`
	}
	if err := Client().JSON(ctx, http.MethodGet, "/api/domains/v1/portfolio", nil, nil, &rows); err != nil {
		return nil, err
	}
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		if r.Domain != "" {
			out = append(out, r.Domain)
		}
	}
	return out, nil
}
