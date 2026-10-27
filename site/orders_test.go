package site

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/cobra"
)

func newOrderCmd() *cobra.Command {
	c := &cobra.Command{Use: "x"}
	c.Flags().String("order-id", "", "")
	return c
}

func TestResolveOrderRungs(t *testing.T) {
	chdir(t, t.TempDir())
	noPrompt(t)
	t.Setenv("HOSTINGER_ORDER_ID", "")

	id, rest, err := ResolveOrder(newOrderCmd(), []string{"42", "box"}, 1, "mail")
	if err != nil || id != "42" || len(rest) != 1 || rest[0] != "box" {
		t.Errorf("positional: %q %v %v", id, rest, err)
	}
	if _, _, err := ResolveOrder(newOrderCmd(), []string{"box"}, 1, "mail"); err != ErrNoOrder {
		t.Errorf("nothing set: err = %v", err)
	}

	dir, _ := os.Getwd()
	if err := Save(dir, &Project{Kind: "website", Domain: "example.com", OrderID: 7}); err != nil {
		t.Fatal(err)
	}
	if id, _, err := ResolveOrder(newOrderCmd(), nil, 0, "hosting"); err != nil || id != "7" {
		t.Errorf("project file for hosting: %q %v", id, err)
	}
	if _, _, err := ResolveOrder(newOrderCmd(), nil, 0, "mail"); err != ErrNoOrder {
		t.Errorf("project file must not feed mail: err = %v", err)
	}

	t.Setenv("HOSTINGER_ORDER_ID", "9")
	if id, _, err := ResolveOrder(newOrderCmd(), nil, 0, "mail"); err != nil || id != "9" {
		t.Errorf("env: %q %v", id, err)
	}
	cmd := newOrderCmd()
	_ = cmd.Flags().Set("order-id", "11")
	if id, _, err := ResolveOrder(cmd, nil, 0, "hosting"); err != nil || id != "11" {
		t.Errorf("flag: %q %v", id, err)
	}
}

func TestListOrdersLabels(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/hosting/v1/orders":
			_ = json.NewEncoder(w).Encode(map[string]any{"data": []map[string]any{{"id": 1, "status": "active", "plan": map[string]string{"name": "Premium"}}}, "meta": map[string]int{"total": 1, "per_page": 100}})
		case "/api/mail/v1/orders":
			_ = json.NewEncoder(w).Encode(map[string]any{"data": []map[string]any{{"id": "abc", "status": "suspended", "domain": map[string]string{"id": "DO1", "name": "example.com"}, "plan": map[string]string{"name": "hostinger_free", "title": "Free Email"}}}, "meta": map[string]int{"total": 1, "per_page": 100}})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	c := api.NewRawClient(srv.URL, nil)
	hosting, err := ListOrders(t.Context(), c, "hosting")
	if err != nil || len(hosting) != 1 || hosting[0].ID != "1" || hosting[0].Label != "Premium (order 1)" {
		t.Errorf("hosting = %+v %v", hosting, err)
	}
	mail, err := ListOrders(t.Context(), c, "mail")
	if err != nil || len(mail) != 1 || mail[0].ID != "abc" || mail[0].Label != "Free Email · example.com (order abc) [suspended]" {
		t.Errorf("mail = %+v %v", mail, err)
	}
	if _, err := ListOrders(t.Context(), c, "ecommerce"); err == nil {
		t.Error("unknown group should fail")
	}
}
