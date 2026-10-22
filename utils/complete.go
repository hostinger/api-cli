package utils

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ArgSource supplies the completions of one positional argument: Fetch lists
// the candidates (it receives the arguments typed so far) and Value names the
// item field offered.
type ArgSource struct {
	Value string
	Fetch func(ctx context.Context, args []string) ([]byte, int, error)
}

// Completion requests must never hold the shell up: they get a short timeout
// and their results are cached briefly per operation and arguments.
var (
	completionTimeout  = 3 * time.Second
	completionCacheTTL = 2 * time.Minute
)

// labelFields are tried in order for the description shown next to a value.
var labelFields = []string{"hostname", "domain", "name", "title", "site_title", "username", "tld", "state", "status"}

// CompleteArgs returns a ValidArgsFunction that completes the next positional
// argument from sources[len(args)]; a nil entry, an API error or a timeout
// yields no completions (and never falls back to file names).
func CompleteArgs(sources ...*ArgSource) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) >= len(sources) || sources[len(args)] == nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		src := sources[len(args)]
		body, ok := readCompletionCache(cmd, args)
		if !ok {
			ctx, cancel := context.WithTimeout(context.Background(), completionTimeout)
			defer cancel()
			b, status, err := src.Fetch(ctx, args)
			if err != nil || status >= 300 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			body = b
			writeCompletionCache(cmd, args, body)
		}
		var out []cobra.Completion
		for _, c := range completionItems(body, src.Value) {
			if strings.HasPrefix(c.value, toComplete) {
				out = append(out, cobra.CompletionWithDesc(c.value, c.label))
			}
		}
		return out, cobra.ShellCompDirectiveNoFileComp
	}
}

type completionItem struct{ value, label string }

// completionItems extracts value/label pairs from a list response: a JSON
// array or a {data: [...]} envelope of objects.
func completionItems(body []byte, field string) []completionItem {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var raw any
	if err := dec.Decode(&raw); err != nil {
		return nil
	}
	if env, ok := raw.(map[string]any); ok {
		raw = env["data"]
	}
	list, _ := raw.([]any)
	var out []completionItem
	for _, it := range list {
		obj, ok := it.(map[string]any)
		if !ok || obj[field] == nil {
			continue
		}
		item := completionItem{value: fmt.Sprintf("%v", obj[field])}
		for _, lf := range labelFields {
			if lf == field {
				continue
			}
			if s, ok := obj[lf].(string); ok && s != "" {
				item.label = s
				break
			}
		}
		out = append(out, item)
	}
	return out
}

// completionCachePath keys the cache by API URL, static token, command and
// arguments (hashed) so different endpoints, tokens and parents never share
// entries.
func completionCachePath(cmd *cobra.Command, args []string) (string, bool) {
	// An OAuth sign-in has no stable identity to key on (another account may
	// sign in), so only static-token results are cached.
	if viper.GetString("api_token") == "" {
		return "", false
	}
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", false
	}
	h := sha256.Sum256([]byte(viper.GetString("api_url") + "\x00" + viper.GetString("api_token") + "\x00" + cmd.CommandPath() + "\x00" + strings.Join(args, "\x00") + "\x00" + fmt.Sprint(len(args))))
	return filepath.Join(dir, cmd.Root().Name(), "completion", hex.EncodeToString(h[:8])+".json"), true
}

func readCompletionCache(cmd *cobra.Command, args []string) ([]byte, bool) {
	path, ok := completionCachePath(cmd, args)
	if !ok {
		return nil, false
	}
	info, err := os.Stat(path)
	if err != nil || time.Since(info.ModTime()) > completionCacheTTL {
		return nil, false
	}
	b, err := os.ReadFile(path)
	return b, err == nil
}

func writeCompletionCache(cmd *cobra.Command, args []string, body []byte) {
	path, ok := completionCachePath(cmd, args)
	if !ok {
		return
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return
	}
	_ = os.WriteFile(path, body, 0o600)
}
