package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// Call performs one API request and returns the raw body and status code.
type Call func(ctx context.Context) ([]byte, int, error)

// maxPages bounds --all so a server that never reports the last page cannot
// loop forever.
const maxPages = 1000

type pageMeta struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
}

// AllPages wraps a paginated list call so it fetches every page (setPage
// selects the page before each request) and returns one {data, meta}
// envelope holding all items. A response that is not a {data: [...]}
// envelope, or an error status, is returned unchanged.
func AllPages(call Call, setPage func(int)) Call {
	return func(ctx context.Context) ([]byte, int, error) {
		var items []json.RawMessage
		for page := 1; page <= maxPages; page++ {
			setPage(page)
			body, status, err := call(ctx)
			if err != nil || status >= 300 {
				return body, status, err
			}
			var env struct {
				Data []json.RawMessage `json:"data"`
				Meta *pageMeta         `json:"meta"`
			}
			if err := json.Unmarshal(body, &env); err != nil || env.Data == nil {
				fmt.Fprintln(promptOut, "Warning: --all: the response is not paginated; showing it as returned.")
				return body, status, nil
			}
			items = append(items, env.Data...)
			if env.Meta == nil && len(env.Data) > 0 {
				fmt.Fprintln(promptOut, "Warning: --all: the response has no page metadata; showing the first page only.")
			}
			if len(env.Data) == 0 || env.Meta == nil || lastPage(env.Meta, page) {
				break
			}
			if page == maxPages {
				fmt.Fprintf(promptOut, "Warning: --all: stopped after %d pages.\n", maxPages)
			}
		}
		if items == nil {
			items = []json.RawMessage{}
		}
		out, err := json.Marshal(map[string]any{
			"data": items,
			"meta": pageMeta{CurrentPage: 1, PerPage: len(items), Total: len(items), LastPage: 1},
		})
		if err != nil {
			return nil, 0, err
		}
		return bytes.TrimSpace(out), 200, nil
	}
}

func lastPage(m *pageMeta, page int) bool {
	if m.LastPage > 0 {
		return page >= m.LastPage
	}
	if m.PerPage <= 0 {
		return true
	}
	return page*m.PerPage >= m.Total
}
