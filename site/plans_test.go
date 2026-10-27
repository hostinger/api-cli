package site

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/api"
)

func TestPlanLabelsSortAndOptions(t *testing.T) {
	sites := []Website{
		{Domain: "a.example.com", OrderID: 12, WebsiteType: "nodejs"},
		{Domain: "b.example.com", OrderID: 11},
		{Domain: "c.example.com", OrderID: 12},
		{Domain: "d.example.com", OrderID: 13},
		{Domain: "e.example.com", OrderID: 99},
	}
	labels := PlanLabels(sites, map[int]string{11: "Single Web Hosting", 12: "Unlimited Web Hosting", 13: "Single Web Hosting"})
	SortByPlan(sites, labels)
	var got []string
	for _, s := range sites {
		got = append(got, PlanOption(s, labels))
	}
	want := "Order 99: e.example.com|Single Web Hosting (order 11): b.example.com|Single Web Hosting (order 13): d.example.com|Unlimited Web Hosting: a.example.com (nodejs)|Unlimited Web Hosting: c.example.com"
	if strings.Join(got, "|") != want {
		t.Errorf("got %q", got)
	}
	// One plan: no grouping.
	one := []Website{{Domain: "x.example.com", OrderID: 11}, {Domain: "y.example.com", OrderID: 11}}
	if labels := PlanLabels(one, map[int]string{11: "Single"}); labels != nil || PlanOption(one[0], labels) != "x.example.com" {
		t.Errorf("one plan: labels=%v", labels)
	}
}

func TestHostingPlanNames(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/hosting/v1/orders":
			json.NewEncoder(w).Encode(map[string]any{
				"data": []map[string]any{
					{"id": 11, "subscription_id": "s1", "plan": map[string]string{"name": "hostinger_starter_v6"}},
					{"id": 12, "subscription_id": "s2", "plan": map[string]string{"name": "hostinger_unlimited_v6"}},
				},
				"meta": map[string]int{"total": 2, "per_page": 100},
			})
		case "/api/billing/v1/subscriptions":
			json.NewEncoder(w).Encode([]map[string]string{{"id": "s1", "name": "Single Web Hosting"}})
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	names, err := HostingPlanNames(context.Background(), api.NewRawClient(srv.URL, nil))
	if err != nil || names[11] != "Single Web Hosting" || names[12] != "hostinger_unlimited_v6" {
		t.Errorf("names=%v err=%v", names, err)
	}
}

func TestNormalizeDomain(t *testing.T) {
	for in, want := range map[string]string{
		"https://darkgray-gull-800972.hostingersite.com/": "darkgray-gull-800972.hostingersite.com",
		"darkgray-gull-800972.hostingersite.com/":         "darkgray-gull-800972.hostingersite.com",
		" HTTP://www.Example.com:8080/a/b?c=d ":           "example.com",
		"example.com.":                                    "example.com",
		"blog.example.com":                                "blog.example.com",
		"":                                                "",
	} {
		if got := NormalizeDomain(in); got != want {
			t.Errorf("%q: got %q, want %q", in, got, want)
		}
	}
}
