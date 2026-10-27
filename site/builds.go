package site

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var ErrNoBuild error = &utils.FlagError{Err: errors.New("no build given: pass the build UUID")}

// ResolveBuild is Resolve plus [build-uuid]: a UUID-shaped arg is the build;
// else the build up last started, else a prompt.
func ResolveBuild(cmd *cobra.Command, args []string, extras int) (username, domain, build string, rest []string, err error) {
	if extras > len(args) {
		extras = len(args)
	}
	lead, rest := args[:len(args)-extras], args[len(args)-extras:]
	if n := len(lead); n > 0 {
		if _, perr := uuid.Parse(lead[n-1]); perr == nil {
			build, lead = lead[n-1], lead[:n-1]
		} else if n == 3 || (n == 2 && !strings.Contains(lead[1], ".")) {
			// Where the build goes but no domain can: a mistyped build, not a website.
			return "", "", "", nil, utils.FlagErrorf("invalid build UUID %q", lead[n-1])
		}
	}
	username, domain, _, err = Resolve(cmd, lead, 0)
	if err != nil {
		return "", "", "", nil, err
	}
	if build != "" {
		return username, domain, build, rest, nil
	}
	if p, _, ferr := Find("."); ferr == nil && p != nil && p.Build != nil && strings.EqualFold(p.Domain, domain) {
		if _, perr := uuid.Parse(p.Build.UUID); perr == nil {
			return username, domain, p.Build.UUID, rest, nil
		}
	}
	if !Interactive() {
		return "", "", "", nil, ErrNoBuild
	}
	ctx := context.Background()
	if cmd != nil && cmd.Context() != nil {
		ctx = cmd.Context()
	}
	build, err = chooseBuild(ctx, username, domain)
	if err != nil {
		return "", "", "", nil, err
	}
	return username, domain, build, rest, nil
}

func chooseBuild(ctx context.Context, username, domain string) (string, error) {
	var resp struct {
		Data []struct {
			UUID      string `json:"uuid"`
			State     string `json:"state"`
			CreatedAt string `json:"created_at"`
		} `json:"data"`
	}
	path := "/api/hosting/v1/accounts/" + url.PathEscape(username) + "/websites/" + url.PathEscape(domain) + "/nodejs/builds"
	if err := Client().JSON(ctx, http.MethodGet, path, nil, nil, &resp); err != nil {
		return "", err
	}
	if len(resp.Data) == 0 {
		return "", fmt.Errorf("%s has no Node.js builds yet: run `hostinger up` to start one", domain)
	}
	if len(resp.Data) == 1 {
		return resp.Data[0].UUID, nil
	}
	options := make([]string, len(resp.Data))
	for i, b := range resp.Data {
		options[i] = fmt.Sprintf("%s  %s  %s", b.UUID, b.State, b.CreatedAt)
	}
	picked, err := pick(pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(len(options) > 8).WithMaxHeight(SelectHeight(len(options))), "Which build?")
	if err != nil {
		return "", err
	}
	for i, o := range options {
		if o == picked {
			return resp.Data[i].UUID, nil
		}
	}
	return "", errors.New("no build selected")
}
